<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#basics">Basics</a>
      <ol>
        <li><a href="#grouping-data">Grouping Data</a></li>
      </ol>
    </li>
    <li><a href="#concurrency">Concurrency</a>
    <li><a href="#channels">Channels</a>
  </ol>
</details>

&nbsp;

## About The Project

- Learn How To Code: Google's Go (golang) Programming Language
- The Ultimate Comprehensive Course - Perfect for Both Beginners and Experienced Developers

&nbsp;

---

&nbsp;

## Basics

- [The Go Blog - Using Go Modules](https://go.dev/blog/using-go-modules)
- Reference Folders: go_modules, iota, control_flow, structs, functions & application

```sh
# Creating a new module
go mod init example.com/username/repo

# Adding a dependency
go get rsc.io/quote

# Upgrading dependencies
go list -m all
go get golang.org/x/text
go get rsc.io/sampler
```

- [The Go Blog - Get familiar with workspaces](https://go.dev/blog/get-familiar-with-workspaces)

> <b>Daniel: </b>Go has two organization tools: modules and packages.
>
> <b>A package is a group of Go files in the same directory.</b> All Go files in the same directory must have the same package statement at the top of the file (one exception in tests), and they can all access all variables, constants, types, and functions from each other. Go files in one package must use an import statement to access values from other packages.
>
> <b>A module is a group of packages.</b> It is defined by a go.mod file, and all go packages in the current directory and sub-directories are a part of that module, unless they have their own go.mod file. A module has a name that looks like an import path. To import a package in another module, you must have the other module listed in your module (dependency management). When you use an import statement to get a package, you include the name of the module in the statement.
>
> For example, I have a module named github.com/myname/awesomemodule and it has a sub-folder named stringtools. To import stringtools, I use import "github.com/myname/awesomemodule/stringtools" If the code attempting to import stringtools is not a part of awesomemodule, I need to add awesomemodule to my go.mod, probably by using go get github.com/myname/awesomemodule from the command line while in one of the directories in my module.

&nbsp;

---

&nbsp;

- Control flow
  - sequence
  - loop; iterative
  - conditional
- [Short variable declarations - declare and assign](https://go.dev/ref/spec#Short_variable_declarations)
- [Types - composite types](https://go.dev/ref/spec#Type_declarations)
- [Conversions](https://go.dev/ref/spec#Conversions)

> <b>Oren: </b>type hotdog int
>
> <b>Todd: </b>The main idea here is that VALUES are of a certain TYPE. You can have the value "TODD" which is of type STRING. You can have the value 42 which is of type INT. When we create our own types, we can specify the UNDERLYING type. This allows us to create code that is self-documenting, among other things
>
> So when you see <code>type hotdog int</code> or <code>type person struct</code>, we have create the TYPE hotdog and person. We can then create VALUES of TYPE hotdog or person. So in our program we might have <code>var x hotdog = 42</code> which would tell us we have 42 hotdogs.
>
> THE MAIN USE of creating our own types, however, is not for something trivial like "hotdog". It is for creating SEMANTICS in our code and for creating data structures, like type person, for aggregating (also known as, gathering, composing) data together. :)

&nbsp;

---

&nbsp;

- [Go Spec - Iota](https://go.dev/ref/spec#Iota)
- [Wiki - Extended Backus–Naur form](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form)

&nbsp;

### Grouping Data

- <b>array</b>
  - a numbered sequence of elements of a single type
  - does not change in size
  - used for Go internals; generally not recommended for your code
- <b>slice</b>
  - built on top of an array
  - changes in size
  - holds values of the same type
  - has a length and a capacity
  - Slices are built on top of arrays. A slice is dynamic in that it will grow in size. The underlying array, however, does not grow in size. When we create a slice, we can use the built in function make to specify how large our slice should be and also how large the underlying array should be. This can enhance performance a little bit.
- <b>map</b>
  - key/value storage
  - an unordered group of elements of one type, called the element type, indexed
    by a set of unique keys of another type, called the key type.
- <b>struct</b>
  - a data structure
  - a composite type
  - allows us to collect values of different types together

&nbsp;

- [Go Standard Library](https://pkg.go.dev/std)
- [Go Forum - Are arrays passed by value or passed by reference in GO?](https://forum.golangbridge.org/t/are-arrays-passed-by-value-or-passed-by-reference-in-go/19255)

&nbsp;

---

&nbsp;

> <b>Matteo: </b>I don't understand what you mean with Overriding: "promotions"
>
> <b>Jared: </b>In a previous lesson, he mentioned that "promotion" means you can access the fields of the "inner" type directly, which equates to them being promoted. For example, in his example he had created a type "person" which he then used as an inner type in another type called "secret agent". To access the fields in "person", he didn't have to go through "person", like "s1.person.first". He could access the fields directly, such as "s1.first". I know it's been a year since you asked the question, but it helps me to remember better if I type out an answer :)

&nbsp;

---

&nbsp;

- <code>func (r receiver) identifier(parameters) (return(s)) {...}</code>
- we define our func with parameters (if any)
- we call our func and pass in arguments (in any)
- <b>Everything in Go is PASS BY VALUE</b>
- A <b>variadic parameter</b> is a func which takes an unlimited number of arguments.
- A <b>method</b> is nothing more than a FUNC attached to a TYPE. When you attach a func to a
  type it is a method of that type. You attach a func to a type with a RECEIVER.
- <b>Interfaces & polymorphism</b>
  - In Go, values can be of more than one type.
  - An interface allows a value to be of more than one type. We create an interface using this syntax: “keyword identifier type” so for an interface it would be: “type human interface”
  - We then define which method(s) a type must have to implement that interface. If a TYPE has the required methods, which could be none (the empty interface denoted by interface{}), then that TYPE implicitly implements the interface and is also of that interface type.
  - In Go, values can be of more than one type.

&nbsp;

> <b>Dharlequin: </b>What is the actual idea/concept behind using defer?

> <b>Daniel: </b>Defer is useful for two circumstances. First is making sure a function is run before the surroundings exit, no matter which return statement is used to do the exiting. Go conventions prefers designing your conditions to exit a function early, instead of other languages (mostly c/c++) method of only using a single return at the end of the function. Which makes defer nice since you don't need to worry about the cleanup at every return.
>
> Defer is also makes sure it runs when a panic occurs. So if something in your function ends up panicking, but it is then recovered further up the call stack, the defer will still have run to get the cleanup done. In this way, you can consider defer similar to other languages concept of try/finally.

&nbsp;

---

&nbsp;

> <b>Abhishek Kumar: </b>what is the purpose of anonymous function? In what kind of scenario can I use them?

> <b>Todd: </b>You might want to "encapsulate" a variable; narrow the scope of some variables; so you could put it all in a function. You also might want to "abstract" out some code; get it modularized so that you could say "this piece does X, this piece does Y, this piece does Z" where each piece was an anonymous func. You also might want to launch a goroutine so you could do this: <code>go func(){<code here>}()</code> Some of the uses!

&nbsp;

---

&nbsp;

- <b>Callback</b>
  - passing a func as an argument
  - functional programming not something that is recommended in go, however, it is good to be aware of callbacks
  - idiomatic go: write clear, simple, readable code
- <b>Closure</b>
  - passing a func as an argument
  - [functional programming](https://stackoverflow.com/questions/66839360/does-go-support-functional-programming) not something that is recommended in go, however, it is good to be aware of callbacks
  - idiomatic go: write clear, simple, readable code
- <b>Pointers</b>
  - All values are stored in memory. Every location in memory has an address. A pointer is a
    memory address.
  - <code>&</code> gives you the address
  - <code>\*</code> gives you the value stored in the address
  - <code>\*int</code>

&nbsp;

---

&nbsp;

- [Go Standard Library encoding/json#Marshal](https://pkg.go.dev/encoding/json#Marshal)
- [JSON to GO](https://mholt.github.io/json-to-go/)
- [JSON encoding vs marshalling](https://stackoverflow.com/questions/33061117/in-golang-what-is-the-difference-between-json-encoding-and-marshalling)

> <b>Jamal: </b>JSON & Marshaling
>
> <b>What is JSON?</b>
> First JSON(JavaScript Object Notation) everyone should know this, it's just keyvalue pairs in JS and it's easier to transmit than XML. It's a human readable, machine transferable and generally the preferred way to to send and receive data via REST APIs. It's not the most efficient way but it's the web-developer preferred way.
>
> <b>What does 'Marshaling' mean?</b>
>
> marshaling is the process of transforming the memory representation of an object to a data format suitable for storage or transmission, and it is typically used when data must be moved between different parts of a computer program or from one program to another.
>
> The inverse of marshaling is called unmarshalling or demarshalling.
>
> <b>As it relates to Golang and this section</b>
>
> This Marshaling and Unmarshalling is Golang trying to convert struct into JSON objects and JSON objects into Golang structs. This section of the course is about how you can transfer to JSON and from JSON(string literal byte slice) back into Golang struct.
>
> Remember Golang is a web backend language, eventually we'll learn how to get JSON objects via HTTP request. While we don't yet know how to do that, Tod is preparing us for that future where we'll have put into a byte slice a JSON object as a string literal. Or the reverse take a JSON object that is a byte slice and convert it to a Golang struct.
>
> Recall that in Golang when you put a string in backtick <code>``</code>, it is treated as raw string literal. Meaning that it is read exactly as is a UTF-8 string with no escape characters only runes(utf-8 characters). Vs the <code>""</code> double quotes which allows for escaped characters.
>
> In quotes <code>""</code> you need to escape new lines, tabs and other characters(in Golang these are called verbs) that do not need to be escaped in backticks <code>``</code>. If you put a line break(\n) in a backtick string, it is interpreted as a <code>'\n'</code> meaning, it's not converted into the actual line break but read as the full characters contained within the ticks.
>
> So now we have this JSON object that we got from HTTP somehow. How do we get it into a struct or how do we get our nice Golang struct into JSON? That is what the JSON package is for.
>
> <b>Marshal function in JSON Package</b>
>
> When we want to convert a Golang struct into a JSON object, we use the json.Marshal. Marshal is Golang way of saying "encode/convert to JSON Object". Because Golang is a strictly typed language and JSON is a dynamically typed language. A few things need to be known while constructing your struct for JSON transfer.
>
> <b>Exposed vs not Exposed fields</b>
>
> As with all structs in Go, it’s important to remember that only fields with a capital first letter are visible to external programs/packages like the JSON Marshaller. Meaning that if you don't capitalize the first letter it won't get exposed when you convert it to JSON using the json.Marshall() function. This is because you cannot access that field outside of the struct, remember if it's not a capital first letter it cannot be accessed outside of the struct.
>
> <b>Meaning of Interface{}</b>
>
> If you’re not familiar an empty <code>interface{}</code> is a way of defining a variable in Go as “this could be anything”. At runtime Go will then allocate the appropriate memory to fit whatever you decide to store in it.
>
> More detail on interface{}
>
> In Go, every type creates a empty interface (interface{}) implicitly (just by existing). A interface is just that, a way to interface with a type! Recall when we were writing explicit interfaces? We had to explicitly say that if a type contains these methods then it is this type.

```go
type Human interface{
  speak()
  walk()
  sleep()
}
```

> This is a example of a interface, something is considered a human by what actions it can do. In programming this is called duck typing, by specifying the actions something can do you specify its type.
>
> Now to explain the Empty interface,
>
> every type in Go implements an empty interface without you having to do anything. A empty interface is, any type that implements zero methods. By explicit design of the Go lang specifications every type at least implements a zero method interface called a empty interface.
>
> <b>Some insight about interfaces</b>
>
> This early on it may not be apparent why interfaces are really cool, but I'll give you a sneak peak. Interfaces allow you to be future proof and allow for additions to code you can't predict right now. What happens if we have the following type in the future?

```go
type cyborg {}
func (c cyborg) speak(){}
func (c cyborg) walk(){}
func (c cyborg) sleep(){}
```

> Without us having to do anything, we can now use the cyborg type anywhere Human type was used previously without any issues. Interfaces are some really genius level constructs they in my opinion are better than class inheritance in other languages. Also they're used extensively in the standard library to simplify code!
>
> <b>One of the things that will get really confusing later on if you don't understand it now</b> by implementing a Reader/Writer you by default get some functionality without having to do any work yourself. I'll expand on this in the notes to the videos in this section. But just know, that when we define a interface to Writer/Reader we get some cool functionality through the interface system!
>
> <b>JSON Marshall</b>
> The function signature for json.marshal
> <b>func Marshal(v interface{}) ([]byte, error)</b>
>
> it takes a v interface{}, which can be 'any' go type. Basically everything from a struct to a primitive it will take it and try to convert it to a JSON object. It returns two things.
>
> 1.  a slice of byte []byte, containing the literal string that is the JSON object.
> 2.  And error, letting you know if anything went wrong
>
> Typically you give the JSON marshal function a pre-filled struct, or a raw string literal formatted to JSON
> <b>How are go types represented in JSON by the json Marshaller?</b>

```go
bool for JSON booleans,
float64 for JSON numbers,
string for JSON strings,
nil for JSON null.
```

> Only data that can be represented as JSON will be encoded(converted) by the json.marshal() function.
>
> Only the exported (public) fields of a struct will be present in the JSON output. A field with a json: tag is stored with its tag name instead of its variable name. Pointers will be encoded as the values they point to, or null if the pointer is nil.
>
> Important: When you tag a struct field with a 'json: tag' you are telling the Marshall package how to interpret the Go Struct. Meaning you can tag something in Json to have a different name than what you are storing in the Go struct. This is really powerful, because you have the ability to define what the json field will be stored as name wise, which can help simplify things to a great deal!
> <b>Unmarshal function in the JSON package</b>
>
> When we want to convert a JSON Object into a Golang struct, we use the json.Unmarshal. Unmarshal is Golang way of saying "parse this JSON object into a valid Golang struct".
> <b>The function signature for unmarshal</b>
>
> func Unmarshal(data []byte, v interface{}) error
>
> unmarshal takes the following parameters:
>
> 1. a slice of bytes (This a raw string, this is the JSON object that you want to parse)
> 2. A pointer to a struct to parse the JSON into
>
> Unmarshal returns,
>
> A error if anything went wrong with parsing.
>
> <b>How does Unmarshal decide which fields to try and parse?</b>
>
> For any key found in the JSON, Unmarshal will try to match it to a key found in the struct with the following logic.
>
> For explanation sakes, I'm using "FieldName" to represent any member of a struct.
>
> 1. It will first look for an exported(field member with a capital letter) with a tag json:"FieldName"
> 2. A exported (field member with a capital letter) with the name FieldName
> 3. Any exported field name, that matches the field name if case-sensitivity is not an issue, e,g fIeLdNaMe, FIELDNAME, fieldname.
>
> <b>ONLY FIELDS FOUNDS IN THE destination type(struct) will be decoded.</b>
>
> Only when a field is found in the destination struct will it be decoded, meaning if there is a field in the JSON that isn't in the destination it will be ignored.
>
> This is useful when you wish to pick only a few specific fields. In particular, any fields not exported in the destination struct will be unaffected.
>
> <b>The escaped character %+V</b>
> The escaped character %+V, in the printf statement in this video does the following. If it's a struct it will print the value of that structure with %V, when you have the +, %+V it will print the members of the struct in the print statement.
> <b>Relationship between strings and byte slices</b>
> Recall that a string in Go is just a sequence of bytes. A byte is just an alias(type byte uint8) for a a uint8. So the underpinnings of a string is a sequence of bytes This is why we can easily do a conversion on a sequence of bytes and turn it into a string and vice versa. So when the marshal function returns a sequence of bytes, remember it's just a raw UTF-8 string.

&nbsp;

---

&nbsp;

> <b>Jerry: </b>Stable and Unstable sort
>
> The stability of a sorting algo. refers to preserving the order of duplicate values.
> in a stable sort by age:
>
> <code>{"James",32},{"Joe",20},{"Mark",20}</code> becomes <code>{"Joe",20},{"Mark",20},{"James",32}</code>
>
> Joe and Mark keep their relative position
> In an unstable sort, this can happen:
>
> <code>{"James",32},{"Joe",20},{"Mark",20}</code> becomes <code>{"Mark",20},{"Joe",20},{"James",32}</code>
>
> Mark and Joe switched positions, but the list is still sorted.

&nbsp;

---

&nbsp;

## Concurrency

- In programming, concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations.
- <b>Concurrency</b> is about dealing with lots of things at once.
- <b>Parallelism</b> is about doing lots of things at once.
- A <b>WaitGroup</b> waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished. Writing concurrent code is super easy: all we do is put “go” in front of a function or method call.
- [Go Standard Library sync#example-WaitGroup](https://pkg.go.dev/sync@go1.18.5#example-WaitGroup)
- [Go spec - Method sets](https://go.dev/ref/spec#Method_sets)
  - [Go Tour - methods/6](https://go.dev/tour/methods/6)
  - [Go Tour - methods/7](https://go.dev/tour/methods/7)
- A <b>“mutex”</b> is a mutual exclusion lock. Mutexes allow us to lock our code so that only one goroutine can access that locked chunk of code at a time.
- [Go Standard Library sync#Mutex](https://pkg.go.dev/sync@go1.18.5#Mutex)
- [Go Standard Library sync/atomic](https://pkg.go.dev/sync/atomic)

```sh
# data race(s)
go run --race todd_mcleod/exercises/main.go

# the go scheduler determines which routines run when
# this is not in your control (mostly)
# you are giving routines to the go scheduler
# it then optimizes and runs them as it sees best
```

&nbsp;

---

&nbsp;

## Channels

- [Concurrency patterns in Golang: WaitGroups and Goroutines](https://blog.logrocket.com/concurrency-patterns-golang-waitgroups-goroutines/)
- [The Go Blog - Go Concurrency Patterns: Context](https://go.dev/blog/context)
- channels allow us to pass values between goroutines
- send means send
  - `S` is after `R` so the arrow goes after chan `make(chan <- int) `
- receive means receive
  - `R` comes before `S` so the arrow goes before chan `make(<-chan int)`
- send & receive (bidirectional)
- **Range:** Range stops reading from a channel when the channel is closed
- **Select:** Select statements pull the value from whatever channel has a value ready to be pulled
- **Fan In:** Taking values from many channels, and putting those values onto one channel
- **Fan Out:** Taking some work and putting the chunks of work onto many goroutines

&nbsp;

---

&nbsp;

> <b>Aman: </b>What does blocking really do?

> <b>Todd: </b>BLOCKING means that you have two things waiting on each other in order for goroutine X to proceed, goroutine X needs goroutine Y to complete AND SIMULTANEOUSLY in order for goroutine Y to proceed, goroutine Y needs goroutine X to complete AND SO THEY ARE BOTH WAITING ON EACH OTHER blocked.

&nbsp;

---

&nbsp;

> <b>Jean-Claude: </b>I don't understand why with a normal channel (none-buffer) and without goroutine we get a "all goroutines are asleep - deadlock" error, while with a buffer-channel we don't have this error.

> <b>Trina: </b>see this from the go lang specs
>
> A new, initialized channel value can be made using the built-in function `make`, which takes the channel type and an optional capacity as arguments:
>
> `make(chan,int)`
>
> The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives). A `nil` channel is never ready for communication.
>
> I was confused as well, so I read the go lang specs

&nbsp;

---

&nbsp;

> <b>Mayur: </b>What is the use of send only and receive only channel?
>
> In send only channel, we can ONLY put data into it, so no body can consume/receive this data. So what is use of send only channel? Same goes for receive only channel

> <b>Todd: </b>Provides clarity and self-documentation.

> <b>Vyacheslav: </b>IMHO a possible use case for send\receive only channels is when you need to pass a non directional channel as an argument to some function and want to make sure that the function will perform only allowed operation on the channel (only send or receive). This way you have a self-documented code that tells you about expected actions inside the function and prevents the function from abusing the channel.

&nbsp;

---

&nbsp;

> <b>Milos: </b>Accidental discovery: Another way to range over channels
> You have something like this in your code for the channel ingestor:

```go
for v := range c {
  fmt.Println(v)
}
```

I accidentally found that this works as well:

```go
for range c {
    fmt.Println(<-c)
}
```

> Code works exactly the same regardless of which of the two for loops I use. Not sure if this is safe or any less idiomatic then the for loop used in the exercise, but i'd like to hear other people's opinions.
>
> <b>IMPORTANT: </b>upon further review, this method is _not_ safe. You will get only 50% of the data via Println, because an incoming data chunk gets "ingested" by the range builtin. Sorry for posting the original comment without testing it more!

&nbsp;

---

&nbsp;
