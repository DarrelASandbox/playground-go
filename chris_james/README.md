# TOC

- [TOC](#toc)
- [shell](#shell)
- [Structs](#structs)
- [Pointers](#pointers)
- [nil](#nil)
- [Errors](#errors)
- [map](#map)
- [Dependency Injection](#dependency-injection)
- [Mocking](#mocking)
  - [Guidelines](#guidelines)
- [goroutines](#goroutines)
- [reflection](#reflection)
- [sync](#sync)

# shell

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

cd 10-concurrency
go test -bench=.

go test -race

go vet
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

People normally get in to a bad state when they don't listen to their tests and are not respecting the refactoring stage.

If your mocking code is becoming complicated or you are having to mock out lots of things to test something, you should listen to that bad feeling and think about your code. Usually it is a sign of

- The thing you are testing is having to do too many things (because it has too many dependencies to mock)
  - Break the module apart so it does less
- Its dependencies are too fine-grained
  - Think about how you can consolidate some of these dependencies into one meaningful module
- Your test is too concerned with implementation details
  - Favour testing expected behavior rather than the implementation

Normally a lot of mocking points to bad abstraction in your code.

What people see here is a weakness in TDD but it is actually a strength, more often than not poor test code is a result of bad design or put more nicely, well-designed code is easy to test.

## Guidelines

The definition of refactoring is that the code changes but the behavior stays the same. If you have decided to do some refactoring in theory you should be able to make the commit without any test changes. So when writing a test ask yourself

- Am I testing the behavior I want, or the implementation details?
- If I were to refactor this code, would I have to make lots of changes to the tests?

Although Go lets you test private functions, I would avoid it as private functions are implementation detail to support public behavior. Test the public behavior. Sandi Metz describes private functions as being "less stable" and you don't want to couple your tests to them.

I feel like if a test is working with more than 3 mocks then it is a red flag - time for a rethink on the design

Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful but that means a tighter coupling between your test code and the implementation. Be sure you actually care about these details if you're going to spy on them

In collaborative projects there is value in auto-generating mocks. In a team, a mock generation tool codifies consistency around the test doubles. This will avoid inconsistently written test doubles which can translate to inconsistently written tests.

You should only use a mock generator that generates test doubles against an interface. Any tool that overly dictates how tests are written, or that use lots of 'magic', can get in the sea.

Without mocking important areas of your code will be untested. In our case we would not be able to test that our code paused between each print but there are countless other examples. Calling a service that can fail? Wanting to test your system in a particular state? It is very hard to test these scenarios without mocking.

Without mocks you may have to set up databases and other third parties things just to test simple business rules. You're likely to have slow tests, resulting in slow feedback loops.

By having to spin up a database or a webservice to test something you're likely to have fragile tests due to the unreliability of such services.

Once a developer learns about mocking it becomes very easy to over-test every single facet of a system in terms of the way it works rather than what it does. Always be mindful about the **value of your tests** and what impact they would have in future refactoring.

# goroutines

Channels are a Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes.

goroutines, the basic unit of concurrency in Go, which let us manage more than one website check request.

anonymous functions, which we used to start each of the concurrent processes that check websites.

channels, to help organize and control the communication between the different processes, allowing us to avoid a race condition bug.

the [race-detector](https://go.dev/blog/race-detector) which helped us debug problems with concurrent code

# reflection

- [reflection](https://go.dev/blog/laws-of-reflection)
- [proposal: spec: type inferred composite literals #12854](https://github.com/golang/go/issues/12854)
- [laws of reflection](https://go.dev/blog/laws-of-reflection)

# sync

- I've seen other examples where the sync.Mutex is embedded into the struct.

```go
type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}
```

Sometimes people forget that embedding types means the methods of that type becomes part of the public interface; and you often will not want that. Remember that we should be very careful with our public APIs, the moment we make something public is the moment other code can couple themselves to it. We always want to avoid unnecessary coupling.

Exposing `Lock` and `Unlock` is at best confusing but at worst potentially very harmful to your software if callers of your type start calling these methods.
