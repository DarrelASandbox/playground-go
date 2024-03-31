- [quii](https://github.com/quii)

```sh
go test
godoc -http=:6060 # http://localhost:6060/pkg/

# Benchmarking
# Benchmarks are run sequentially.
go test -bench=.

go test -cover

cd 05-structs-methods-interfaces
go test -run TestArea/Rectangle

# Go linter
go install github.com/kisielk/errcheck@latest
errcheck .
```

It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. Every test has a cost.

# Structs

Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer

Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)

Adding methods so you can add functionality to your data types and so you can implement interfaces

[Table driven tests](https://go.dev/wiki/TableDrivenTests) to make your assertions clearer and your test suites easier to extend & maintain

At some point you may wish to use structs to manage state, exposing methods to let users change the state in a way that you can control.

Pointers to struct have their own name: struct pointers and they are [automatically dereferenced](https://go.dev/ref/spec#Method_values)

# Pointers

Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.

The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

# nil

Pointers can be nil

When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.

Useful for when you want to describe a value that could be missing

# Errors

[Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

# map

[If a map isn’t a reference variable, what is it?](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it)

An interesting property of maps is that you can modify them without passing as an address to it (e.g `&myMap`)

This may make them feel like a "reference type" which they are not
(A map value is a pointer to a `runtime.hmap` structure.)

So when you pass a map to a function/method, you are indeed copying it, but just the pointer part,
not the underlying data structure that contains the data.

A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.

Therefore, you should never initialize an empty map variable:
`var m map[string]string`

Instead, you can initialize an empty map like we were doing above, or use the make keyword to create a map for you:

```go
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```

Both approaches create an empty `hash map` and point `dictionary` at it. Which ensures that you will never get a runtime panic.

# Dependency Injection

- You don't need a framework
- It does not overcomplicate your design
- It facilitates testing
- It allows you to write great, general-purpose functions

Motivated by our tests we refactored the code so we could control where the data was written by injecting a dependency which allowed us to:

- **Test our code**: If you can't test a function easily, it's usually because of dependencies hard-wired into a function or global state. If you have a global database connection pool for instance that is used by some kind of service layer, it is likely going to be difficult to test and they will be slow to run. DI will motivate you to inject in a database dependency (via an interface) which you can then mock out with something you can control in your tests.
- **Separate our concerns**, decoupling where the data goes from how to generate it. If you ever feel like a method/function has too many responsibilities (generating data and writing to a db? handling HTTP requests and doing domain level logic?) DI is probably going to be the tool you need.
- **Allow our code to be re-used in different contexts**: The first "new" context our code can be used in is inside tests. But further on if someone wants to try something new with your function they can inject their own dependencies.

# Mocking

- Every forward-thinking post about software development emphasizes the importance of quick feedback loops.
- Slow tests ruin developer productivity.
- If we can mock `time.Sleep` we can use dependency injection to use it instead of a "real" `time.Sleep` and then we can **spy on the calls** to make assertions on them.
