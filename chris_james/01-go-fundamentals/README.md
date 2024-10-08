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
  - [When to use locks over channels and goroutines?](#when-to-use-locks-over-channels-and-goroutines)
  - [Don't use embedding because it's convenient](#dont-use-embedding-because-its-convenient)
- [context](#context)
  - [context.Value](#contextvalue)
- [Property Based Tests](#property-based-tests)
  - [Domain of Roman Numerals](#domain-of-roman-numerals)
- [Maths](#maths)
  - [An SVG of a clock](#an-svg-of-a-clock)
  - [Floating Point Math](#floating-point-math)
  - [A note on dividing by zero](#a-note-on-dividing-by-zero)
  - [The Most Valuable Test](#the-most-valuable-test)
- [Reading Blogposts](#reading-blogposts)
  - [File system abstractions introduced in Go 1.16](#file-system-abstractions-introduced-in-go-116)
  - [Additional Implementation Details](#additional-implementation-details)
  - [Further Reading](#further-reading)
- [Templating](#templating)
  - [Embed](#embed)
  - [Approval Tests](#approval-tests)
  - [`template.FuncMap`](#templatefuncmap)
  - [Separating Concerns](#separating-concerns)
    - [Add a method to Post and then call that in the template](#add-a-method-to-post-and-then-call-that-in-the-template)
    - [Create a dedicated view model type, such as `PostViewModel` with exactly the data we need](#create-a-dedicated-view-model-type-such-as-postviewmodel-with-exactly-the-data-we-need)
- [Generics And Arrays](#generics-and-arrays)
  - [Functional Programming](#functional-programming)
  - [Identity Element](#identity-element)

> **Write the test we want to see.** Think about how we'd like to use the code we're going to write from a consumer's point of view.
>
> Focus on what and why, but don't get distracted by how.

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

# Run tests with verbose output
go test -v

# To run all tests in a Go project, including those in subdirectories
go test ./...
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

## When to use locks over channels and goroutines?

- [Go Wiki: Use a sync.Mutex or a channel?](https://go.dev/wiki/MutexOrChannel)
  - Use channels when passing ownership of data
  - Use mutexes for managing state

## Don't use embedding because it's convenient

- Think about the effect embedding has on your public API.
- Do you really want to expose these methods and have people coupling their own code to them?
- With respect to mutexes, this could be potentially disastrous in very unpredictable and weird ways, imagine some nefarious code unlocking a mutex when it shouldn't be; this would cause some very strange bugs that will be hard to track down.

# context

Software often kicks off long-running, resource-intensive processes (often in goroutines). If the action that caused this gets cancelled or fails for some reason you need to stop these processes in a consistent way through your application.

If you don't manage this your snappy Go application that you're so proud of could start having difficult to debug performance problems.

We'll use the package `context` to help us manage long-running processes.

## context.Value

Some engineers have advocated passing values through `context` as it feels convenient.

Convenience is often the cause of bad code.

The problem with `context.Values` is that it's just an untyped map so you have no type-safety and you have to handle it not actually containing your value. You have to create a coupling of map keys from one module to another and if someone changes something things start breaking.

In short, **if a function needs some values, put them as typed parameters rather than trying to fetch them from**
`context.Value`. This makes it statically checked and documented for everyone to see.

But...

On other hand, it can be helpful to include information that is orthogonal to a request in a context, such as a trace id. Potentially this information would not be needed by every function in your call-stack and would make your functional signatures very messy.

[Jack Lindamood says Context.Value should inform, not control](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39)

[Context should go away for Go 2](https://faiface.github.io/post/context-should-go-away-go2/)

[Go Concurrency Patterns: Context](https://go.dev/blog/context)

# Property Based Tests

Property based tests help you do this by throwing random data at your code and verifying the rules you describe always hold true. A lot of people think property based tests are mainly about random data but they would be mistaken. The real challenge about property based tests is having a good understanding of your domain so you can write these properties.

## Domain of Roman Numerals

1. Can't have more than 3 consecutive symbols
2. Only I (1), X (10) and C (100) can be "subtractors"
3. Taking the result of `ConvertToRoman(N)` and passing it to `ConvertToArabic` should return us `N`

# Maths

## An SVG of a clock

- [XML to Go](https://xml-to-go.github.io/)

```svg
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">

  <!-- bezel -->
  <circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>

  <!-- hour hand -->
  <line x1="150" y1="150" x2="114.150000" y2="132.260000"
        style="fill:none;stroke:#000;stroke-width:7px;"/>

  <!-- minute hand -->
  <line x1="150" y1="150" x2="101.290000" y2="99.730000"
        style="fill:none;stroke:#000;stroke-width:7px;"/>

  <!-- second hand -->
  <line x1="150" y1="150" x2="77.190000" y2="202.900000"
        style="fill:none;stroke:#f00;stroke-width:3px;"/>
</svg>
```

## Floating Point Math

- [0.30000000000000004](https://0.30000000000000004.com/)

- There are two ways around this:
  - Live with it
  - Refactor our function by refactoring our equation

Now (1) may not seem all that appealing, but it's often the only way to make floating point equality work. Being inaccurate by some infinitesimal fraction is frankly not going to matter for the purposes of drawing a clockface, so we could write a function that defines a 'close enough' equality for our angles. But there's a simple way we can get the accuracy back: we rearrange the equation so that we're no longer dividing down and then multiplying up. We can do it all by just dividing.

So instead of `numberOfSeconds * π / 30` we can write `π / (30 / numberOfSeconds)`

## A note on dividing by zero

Computers often don't like dividing by zero because infinity is a bit strange.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(secondsinradians())
}

func zero() float64 {
	return 0.0
}

func secondsinradians() float64 {
	return (math.Pi / (30 / (float64(zero()))))
}
```

## The Most Valuable Test

The most sophisticated code for handling SVGs in our project is found in our test code, not in the application code. This might feel uncomfortable, prompting us to consider using a text/template, an XML library, or an SVG library to refactor our code. However, the critical aspect is not how we produce the SVG but ensuring that the output is a valid SVG.

The test code needs to understand SVGs thoroughly to ensure our output is correct. While it might seem excessive to invest significant time in these SVG tests, this effort is crucial. These tests, which involve importing an XML library, parsing XML, and refactoring structs, are highly valuable. They ensure our SVG output remains valid regardless of the production code changes.

Tests are not secondary; they are essential and often more enduring than the code they test. Investing time in writing good tests is worthwhile and should never be seen as spending 'too much time.' It's an investment in the quality and reliability of our codebase.

# Reading Blogposts

```go
var posts []blogposts.Post
posts blogposts.NewPostsFromFS("some-folder")
```

- To write a test around this, we'd need some kind of test folder with some example posts in it. There's nothing terribly wrong with this, but you are making some trade-offs:
  - for each test you may need to create new files to test a particular behavior
  - some behavior will be challenging to test, such as failing to load files
  - the tests will run a little slower because they will need to access the file system
- We're also unnecessarily coupling ourselves to a specific implementation of the file system.

## File system abstractions introduced in Go 1.16

> [On the producer side of the interface, the new embed.FS type implements fs.FS, as does zip.Reader. The new os.DirFS function provides an implementation of fs.FS backed by a tree of operating system files.](https://go.dev/doc/go1.16#fs)

If we use this interface, users of our package have a number of options baked-in to the standard library to use. Learning to leverage interfaces defined in Go's standard library (e.g. `io.fs`, `io.Reader`, `io.Writer`), is vital to writing loosely coupled packages. These packages can then be re-used in contexts different to those you imagined, with minimal fuss from your consumers.

```go
var posts []blogposts.Post
posts = blogposts.NewPostsFromFS(someFS)
```

- [How to level up your TDD skills?](https://deniseyu.github.io/leveling-up-tdd/)

## Additional Implementation Details

- Need to handle:
  - when the file's format is not correct
  - the file is not a `.md`
  - what if the order of the metadata fields is different? Should that be allowed? Should we be able to handle it?

## Further Reading

- [A Tour of Go 1.16's io/fs package](https://benjamincongdon.me/blog/2021/01/21/A-Tour-of-Go-116s-iofs-package/)
- [io/fs: add file system interfaces #41190](https://github.com/golang/go/issues/41190)

# Templating

## Embed

> Package embed provides access to files embedded in the running Go program.
> Go source files that import "embed" can use the //go:embed directive to initialize a variable of type string, []byte, or FS with the contents of files read from the package directory or subdirectories at compile time.

Why would we want to use this? Well the alternative is that we can load our templates from a "normal" file system. However this means we'd have to make sure that the templates are in the correct file path wherever we want to use this software. In your job you may have various environments like development, staging and live. For this to work, you'd need to make sure your templates are copied to the correct place.

With embed, the files are included in your Go program when you build it. This means once you've built your program (which you should only do once), the files are always available to you.

What's handy is you can not only embed individual files, but also file systems; and that filesystem implements [`io/fs`](https://pkg.go.dev/io/fs) which means your code doesn't need to care what kind of file system it is working with.

If you wish to use different templates depending on configuration though, you may wish to stick to loading templates from disk in the more conventional way.

## Approval Tests

- [approvals/go-approval-tests](https://github.com/approvals/go-approval-tests)

> ApprovalTests allows for easy testing of larger objects, strings and anything else that can be saved to a file (images, sounds, CSV, etc...)

The idea is similar to "golden" files, or snapshot testing. Rather than awkwardly maintaining strings within a test file, the approval tool can compare the output for you with an "approved" file you created. You then simply copy over the new version if you approve it. Re-run the test and you're back to green.

- [Part 1/3 - Introducing the Gilded Rose kata and writing test cases using Approval Tests, Emily Bache](https://www.youtube.com/watch?v=zyM2Ep28ED8)

## `template.FuncMap`

Before you parse a template you can add a `template.FuncMap` into your template, which allows you to define functions that can be called within your template. In this case we've made a `sanitizeTitle` function which we then call inside our template with `{{sanitizeTitle .Title}}`.

This is a powerful feature, being able to send functions in to your template will allow you to do some very cool things, but, should you? Going back to the principles of Mustache and logic-less templates, why did they advocate for logic-less? **What is wrong with logic in templates?**

As we've shown, in order to test our templates, we've had to introduce a whole different kind of testing.

Imagine you introduce a function into a template which has a few different permutations of behaviour and edge cases, **how will you test it?** With this current design, your only means of testing this logic is by rendering HTML and comparing strings. This is not an easy or sane way of testing logic, and definitely not what you'd want for important business logic.

Even though the approval tests technique has reduced the cost of maintaining these tests, they're still more expensive to maintain than most unit tests you'll write. They're still sensitive to any minor markup changes you might make, it's just we've made it easier to manage. We should still strive to architect our code so we don't have to write many tests around our templates, and try and separate concerns so any logic that doesn't need to live inside our rendering code is properly separated.

What Mustache-influenced templating engines give you is a useful constraint, don't try to circumvent it too often; **don't go against the grain**. Instead, embrace the idea of view models, where you construct specific types that contain the data you need to render, in a way that's convenient for the templating language.

This way, whatever important business logic you use to generate that bag of data can be unit tested separately, away from the messy world of HTML and templating.

## Separating Concerns

### Add a method to Post and then call that in the template

We can call methods in our templating code on the types we send, so we could add a `SanitizedTitle` method to `Post`. This would simplify the template and we could easily unit test this logic separately if we wish. This is probably the easiest solution, although not necessarily the simplest.

A downside to this approach is that this is still view logic. It's not interesting to the rest of the system but it now becomes a part of the API for a core domain object. This kind of approach over time can lead to you creating [God Objects](https://en.wikipedia.org/wiki/God_object).

### Create a dedicated view model type, such as `PostViewModel` with exactly the data we need

```go
type PostViewModel struct {
  Title, SanitizeTitle, Description, Body string
	Tags                                    []string
}
```

A way to keep this clean would be to have a `func NewPostView(p Post) PostView` which would encapsulate the mapping. Callers of our code would have to map from `[]Post` to `[]PostView`, generating the `SanitizedTitle`.

This would keep our rendering code logic-less and is probably the strictest separation of concerns we could do, but the trade-off is a slightly more convoluted process to get our posts rendered.

# Generics And Arrays

## Functional Programming

> In functional programming, fold (also termed reduce, accumulate, aggregate, compress, or inject) refers to a family of higher-order functions that analyze a recursive data structure and through use of a given combining operation, recombine the results of recursively processing its constituent parts, building up a return value. Typically, a fold is presented with a combining function, a top node of a data structure, and possibly some default values to be used under certain conditions. The fold then proceeds to combine elements of the data structure's hierarchy, using the function in a systematic way.

- [Fold (higher-order function)](<https://en.wikipedia.org/wiki/Fold_(higher-order_function)>)

> **Don't conflate easiness, with simplicity.** Doing loops and copy-pasting code is easy, but it's not necessarily simple. For more on simple vs easy, watch [Rich Hickey's masterpiece of a talk - Simple Made Easy](https://www.youtube.com/watch?v=SxdOUGdseq4).

> **Don't conflate unfamiliarity, with complexity.** Fold/reduce may initially sound scary and computer-sciencey but all it really is, is an abstraction over a very common operation. Taking a collection, and combining it into one item. When you step back, you'll realize you probably do this a lot.

## Identity Element

- [Identity Element](https://en.wikipedia.org/wiki/Identity_element)

> In mathematics, an identity element, or neutral element, of a binary operation operating on a set is an element of the set which leaves unchanged every element of the set when the operation is applied.
