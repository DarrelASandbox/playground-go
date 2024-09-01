- [Question \& Answer](#question--answer)
  - [OS Exec](#os-exec)
  - [Error Types](#error-types)
  - [Context-Aware Reader](#context-aware-reader)
  - [HTTP Handlers](#http-handlers)
    - [Given Code](#given-code)
    - [What is a HTTP Handler and what should it do?](#what-is-a-http-handler-and-what-should-it-do)
    - [Separate of concern for given code](#separate-of-concern-for-given-code)

# Question & Answer

## OS Exec

keith6014 asks on [reddit](https://www.reddit.com/r/golang/comments/aaz8ji/testdata_and_function_setup_help/)

> I am executing a command using os/exec.Command() which generated XML data. The command will be executed in a function called GetData().
>
> In order to test GetData(), I have some testdata which I created.
>
> In my \_test.go I have a TestGetData which calls GetData() but that will use os.exec, instead I would like for it to use my testdata.
>
> What is a good way to achieve this? When calling GetData should I have a "test" flag mode so it will read a file ie GetData(mode string)?

- **Tips**:
  - When something is difficult to test, it's often due to the separation of concerns not being quite right
  - Don't add "test modes" into your code, instead use Dependency Injection so that you can model your dependencies and separate concerns.

## Error Types

Pedro on the Gopher Slack asks

> If I’m creating an error like `fmt.Errorf("%s must be foo, got %s", bar, baz)`, is there a way to test equality without comparing the string value?

- **Tips**:
  - If you find yourself testing for multiple error conditions don't fall in to the trap of comparing the error messages.
  - [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)

## Context-Aware Reader

This chapter demonstrates how to test-drive a context aware `io.Reader` as written by Mat Ryer and David Hernandez in [The Pace Dev Blog](https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer.html).

[In a previous chapter](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context) we discussed how we can use context to provide cancellation. This is especially useful if you're performing tasks which may be computationally expensive and you want to be able to stop them.

- What we want to demonstrate is something like
  - Given an `io.Reader` with "ABCDEF", when I send a cancel signal half-way through I when I try to continue to read I get nothing else so all I get is "ABC"
- So rather than reading everything, we could:
  - Supply a fixed-size byte array that doesn't fit all the contents
  - Send a cancel signal
  - Try and read again and this should return an error with 0 bytes read
- Summary:
  - Small interfaces are good and are easily composed
  - When you're trying to augment one thing (e.g io.Reader) with another you usually want to reach for the [delegation pattern](https://en.wikipedia.org/wiki/Delegation_pattern)

> In software engineering, the delegation pattern is an object-oriented design pattern that allows object composition to achieve the same code reuse as inheritance.

- An easy way to start this kind of work is to wrap your delegate and write a test that asserts it behaves how the delegate normally does before you start composing other parts to change behaviour. This will help you to keep things working correctly as you code toward your goal

## HTTP Handlers

[Santosh Kumar tweeted me](https://x.com/sntshk/status/1255559003339284481)

> How do I test a http handler which has mongodb dependency?

### Given Code

```go
func Registration(w http.ResponseWriter, r *http.Request) {
	var res model.ResponseResult
	var user model.User

	w.Header().Set("Content-Type", "application/json")

	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()
	defer r.Body.Close()

	// check if there is proper json body or error
	if err := jsonDecoder.Decode(&user); err != nil {
		res.Error = err.Error()
		// return 400 status codes
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	// Connect to mongodb
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	// Check if username already exists in users datastore, if so, 400
	// else insert user right away
	collection := client.Database("test").Collection("users")
	filter := bson.D{{"username", user.Username}}
	var foundUser model.User
	err = collection.FindOne(context.TODO(), filter).Decode(&foundUser)
	if foundUser.Username == user.Username {
		res.Error = UserExists
		// return 400 status codes
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		res.Error = err.Error()
		// return 400 status codes
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}
	user.Password = string(pass)

	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		res.Error = err.Error()
		// return 400 status codes
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	// return 200
	w.WriteHeader(http.StatusOK)
	res.Result = fmt.Sprintf("%s: %s", UserCreated, insertResult.InsertedID)
	json.NewEncoder(w).Encode(res)
	return
}
```

- List all the things this one function has to do:
  - Write HTTP responses, send headers, status codes, etc.
  - Decode the request's body into a User
  - Connect to a database (and all the details around that)
  - Query the database and applying some business logic depending on the result
  - Generate a password
  - Insert a record

### [What is a HTTP Handler and what should it do?](https://quii.gitbook.io/learn-go-with-tests/questions-and-answers/http-handlers-revisited#what-is-a-http-handler-and-what-should-it-do)

1. Accept a HTTP request, parse and validate it.
2. Call some `ServiceThing` to do `ImportantBusinessLogic` with the data I got from step 1.
3. Send an appropriate `HTTP` response depending on what `ServiceThing` returns.

- When you separate these concerns:
  - Testing handlers becomes a breeze and is focused a small number of concerns.
  - Importantly testing ImportantBusinessLogic no longer has to concern itself with `HTTP`, you can test the business logic cleanly.
  - You can use `ImportantBusinessLogic` in other contexts without having to modify it.
  - If `ImportantBusinessLogic` changes what it does, so long as the interface remains the same you don't have to change your handlers.

### Separate of concern for given code

1. Decode the request's body into a `User`
2. Call a `UserService.Register(user)` (this is our `ServiceThing`)
3. If there's an error act on it (the example always sends a `400 BadRequest` which I don't think is right), I'll just have a catch-all handler of a `500 Internal Server Error` for now. I must stress that returning `500` for all errors makes for a terrible API! Later on we can make the error handling more sophisticated, perhaps with error types.
4. If there's no error, `201 Created` with the ID as the response body (again for terseness/laziness)
