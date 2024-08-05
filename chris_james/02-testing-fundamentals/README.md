- [Introduction To Acceptance Tests](#introduction-to-acceptance-tests)
  - [What are they?](#what-are-they)
  - [Benefits of acceptance tests](#benefits-of-acceptance-tests)
  - [Potential drawbacks vs unit tests](#potential-drawbacks-vs-unit-tests)
  - [High-level steps for the acceptance test](#high-level-steps-for-the-acceptance-test)
  - [`LaunchTestProgram`](#launchtestprogram)
- [Scaling Acceptance Tests](#scaling-acceptance-tests)
  - [shell](#shell)
  - [Anatomy of bad acceptance tests](#anatomy-of-bad-acceptance-tests)
  - [Tight coupling](#tight-coupling)
  - [First system: HTTP API](#first-system-http-api)
    - [Adapter Pattern](#adapter-pattern)
    - [Reflect](#reflect)
    - [Complexity](#complexity)
    - [High-level steps for acceptance test](#high-level-steps-for-acceptance-test)
    - [New functionality](#new-functionality)
  - [Second System: GRPC](#second-system-grpc)
    - [Overview](#overview)
    - [When should I write acceptance tests?](#when-should-i-write-acceptance-tests)
  - [Third System: Web](#third-system-web)
- [Working Without Mocks](#working-without-mocks)
  - [A primer on test doubles](#a-primer-on-test-doubles)
    - [Stubs](#stubs)
    - [Spies](#spies)
    - [Mocks](#mocks)
    - [Fakes](#fakes)
  - [The problem with stubs and mocks](#the-problem-with-stubs-and-mocks)
  - [Evolving software](#evolving-software)
- [ent](#ent)

# Introduction To Acceptance Tests

- [Are you an Elite DevOps performer? Find out with the Four Keys Project](https://cloud.google.com/blog/products/devops-sre/using-the-four-keys-to-measure-your-devops-performance)
- [Kubernetes best practices: terminating with grace](https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-terminating-with-grace)

- What we want to do is listen for `SIGTERM`, and rather than instantly killing the server, we want to:
  - Stop listening to any more requests
  - Allow any in-flight requests to finish
  - Then terminate the process

## What are they?

- Acceptance tests are a kind of "black-box test". They are sometimes referred to as "functional tests". They should exercise the system as a user of the system would.
- The term "black-box" refers to the idea that the test code has no access to the internals of the system, it can only use its public interface and make assertions on the behaviors it observes. This means they can only test the system as a whole.
- This is an advantageous trait because it means the tests exercise the system the same as a user would, it can't use any special workarounds that could make a test pass, but not actually prove what you need to prove. This is similar to the principle of preferring your unit test files to live inside a separate test package, for example, package `mypkg_test` rather than package `mypkg`.

## Benefits of acceptance tests

- When they pass, you know your entire system behaves how you want it to.
- They are more accurate, quicker, and require less effort than manual testing.
- When written well, they act as accurate, verified documentation of your system. It doesn't fall into the trap of documentation that diverges from the real behavior of the system.
- No mocking! It's all real.

## Potential drawbacks vs unit tests

- They are expensive to write.
- They take longer to run.
- They are dependent on the design of the system.
- When they fail, they typically don't give you a root cause, and can be difficult to debug.
- They don't give you feedback on the internal quality of your system. You could write total garbage and still make an acceptance test pass.
- Not all scenarios are practical to exercise due to the black-box nature.

- [The Practical Test Pyramid](https://martinfowler.com/articles/practical-test-pyramid.html)

## High-level steps for the acceptance test

- Build the program
- Run it (and wait for it listen on `8080`)
- Send an HTTP request to the server
- Before the server has a chance to send an HTTP response, send `SIGTERM`
- See if we still get a response

## `LaunchTestProgram`

- building the program
- launching the program
- waiting for it to listen on port `8080`
- providing a `cleanup` function to kill the program and delete it to ensure that when our tests finish, we're left in a clean state
- providing an `interrupt` function to send the program a SIGTERM to let us test the behaviour

# Scaling Acceptance Tests

- [Dave Farley - How to write acceptance tests](https://www.youtube.com/watch?v=JDD5EEJgpHU)
- [Nat Pryce - E2E functional tests that can run in milliseconds](https://www.youtube.com/watch?v=Fk4rCn4YLLU)
- [Growing Object-Oriented Software Guided by Tests](www.growing-object-oriented-software.com)

## shell

```sh
# Run all tests, ignoring cached results
go test -count=1 ./...

# Individual test
go test -v ./cmd/httpserver
go test -v ./cmd/grpcserver
go test -v ./cmd/webserver

# Test the Container Manually
docker build -t greeter-server -f cmd/httpserver/Dockerfile .
docker run -p 8080:8080 greeter-server
curl http://localhost:8080/greet
```

## Anatomy of bad acceptance tests

- Slow to run
- Brittle
- Flaky
- Expensive to maintain, and seem to make changing the software harder than it ought to be
- Can only run in a particular environment, causing slow and poor feedback loops

## Tight coupling

- Think about what prompts acceptance tests to change:
  - An external behavior change. If you want to change what the system does, changing the acceptance test suite seems reasonable, if not desirable.
  - An implementation detail change / refactoring. Ideally, this shouldn't prompt a change, or if it does, a minor one.

## First system: HTTP API

1. A **driver**. In this case, one works with an HTTP system by using an **HTTP client**. This code will know how to work with our API. Drivers translate DSLs into system-specific calls; in our case, the driver will implement the interface specifications define.
2. An **HTTP server** with a greet API
3. A **test**, which is responsible for managing the life-cycle of spinning up the server and then plugging the driver into the specification to run it as a test

### Adapter Pattern

> In software engineering, the adapter pattern is a software design pattern (also known as wrapper, an alternative naming shared with the decorator pattern) that allows the interface of an existing class to be used as another interface.[1] It is often used to make existing classes work with others without modifying their source code.

### Reflect

- Analyze your problem and identify a slight improvement to your system that pushes you in the right direction
- Capture the new essential complexity in a specification
- Follow the compilation errors until the AT runs
- Update your implementation to make the system behave according to the specification
- Refactor

### Complexity

In the context of Go testing, **essential complexity** refers to the inherent challenges and intricacies that arise directly from the problem you're trying to solve. This includes the logic of the tests, the conditions you need to cover, and the actual behavior of the system under test. For example, if you're testing a complex algorithm or a system with many interdependent components, the complexity of ensuring all paths are tested correctly is essential.

**Accidental complexity**, on the other hand, arises from the tools, frameworks, or practices used to implement the tests, rather than from the problem itself. In Go, this could include difficulties with setting up the test environment, managing dependencies, or dealing with limitations or quirks in the testing framework. For instance, if your tests are hard to write or maintain because of the way the test setup is structured, that’s accidental complexity.

The goal in Go testing, as with any testing, is to minimize accidental complexity so that you can focus on the essential complexity—ensuring your code works correctly under all necessary conditions. Go's standard testing package is designed to keep accidental complexity low by providing simple and effective tools for writing tests.

### High-level steps for acceptance test

- Build a docker image
- Wait for it to be listening on some port
- Create a driver that understands how to translate the DSL into system specific calls
- Plug in the driver into the specification

### New functionality

- Shouldn't have to change the specification;
- Should be able to reuse the specification;
- Should be able to reuse the domain code.

## Second System: GRPC

```sh
# Generate the client and server code
cd adapters/grpcserver
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    greet.proto
```

### Overview

- `adapters` have cohesive units of functionality grouped together
- `cmd` holds our applications and corresponding acceptance tests
- Our code is totally decoupled from any accidental complexity

### When should I write acceptance tests?

1. Is this an edge case? I'd prefer to unit test those
2. Is this something that the non-computer people talk about a lot? I would prefer to have a lot of confidence the key thing "really" works, so I'd add an acceptance test
3. Am I describing a user journey, rather than a specific function? Acceptance test
4. Would unit tests give me enough confidence? Sometimes you're taking an existing journey that already has an acceptance test, but you're adding other functionality to deal with different scenarios due to different inputs. In this case, adding another acceptance test adds a cost but brings little value, so I'd prefer some unit tests.

## Third System: Web

- [GopherCon UK 2021: Riya Dattani & Chris James - Acceptance Tests, BDD & GO](https://www.youtube.com/watch?v=ZMWJCk_0WrY)

> Imagine the least technical person that you can think of, who understands the problem-domain, reading your Acceptance Tests. The tests should make sense to that person.

# Working Without Mocks

- Mocks, spies and stubs encourage you to encode assumptions of the behavior of your dependencies ad-hocly in each test.
- These assumptions are usually not validated beyond manual checking, so they threaten your test suite's usefulness.
- Fakes and contracts give us a more sustainable method for creating test doubles with validated assumptions and better reuse than the alternatives.

> In Mocking, we learned how mocks, stubs and spies are useful tools for controlling and inspecting the behaviour of units of code in conjunction with Dependency Injection.
>
> As a project grows, though, these kinds of test doubles can become a maintenance burden, and we should instead look to other design ideas to keep our system easy to reason and test.
>
> Fakes and contracts allow developers to test their systems with more realistic scenarios, improve local development experience with faster and more accurate feedback loops, and manage the complexity of evolving dependencies.

## [A primer on test doubles](https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/working-without-mocks#a-primer-on-test-doubles)

- **Test doubles** is the collective noun for the different ways you can construct dependencies that you can control for a **subject under test (SUT)**, the thing you're testing. Test doubles are often a better alternative than using the real dependency as it can avoid issues like
  - Needing the internet to use an API
  - Avoid latency and other performance issues
  - Unable to exercise non-happy path cases
  - Decoupling your build from another team's.

### Stubs

```go
// Stubs return the same canned data every time they are called
type StubRecipeStore struct {
	recipes []Recipe
	err     error
}

func (s *StubRecipeStore) GetRecipes() ([]Recipe, error) {
	return s.recipes, s.err
}

// AddRecipes omitted for brevity

// in test, we can set up the stub to always return specific recipes, or an error
stubStore := &StubRecipeStore{
	recipes: someRecipes,
}
```

### Spies

```go
// Spies are like stubs but also record how they were called so the test can assert that the SUT calls the dependencies in specific ways.
type SpyRecipeStore struct {
	AddCalls [][]Recipe
	err      error
}

func (s *SpyRecipeStore) AddRecipes(r ...Recipe) error {
	s.AddCalls = append(s.AddCalls, r)
	return s.err
}

// GetRecipes omitted for brevity

// in test
spyStore := &SpyRecipeStore{}
sut := NewThing(spyStore)
sut.DoStuff()

// now we can check the store had the right recipes added by inspection spyStore.AddCalls
```

### Mocks

```go
/*
Mocks are like a superset of the above, but they only respond with specific data to specific invocations.
If the SUT calls the dependencies with the wrong arguments, it'll typically panic.
 */

// set up the mock with expected calls
mockStore := &MockRecipeStore{}
mockStore.WhenCalledWith(someRecipes).Return(someError)

// when the sut uses the dependency, if it doesn't call it with someRecipes, usually mocks will panic
```

### Fakes

```go
/*
Fakes are like a genuine version of the dependency but implemented in a way more suited to fast running,
reliable tests and local development.

Often, your system will have some abstraction around persistence, which will be implemented with a database,
but in your tests, you could use an in-memory fake instead.
 */

type FakeRecipeStore struct {
	recipes []Recipe
}

func (f *FakeRecipeStore) GetRecipes() ([]Recipe, error) {
	return f.recipes, nil
}

func (f *FakeRecipeStore) AddRecipes(r ...Recipe) error {
	f.recipes = append(f.recipes, r...)
	return nil
}
```

- Fakes are useful because:
  - Their statefulness is useful for tests involving multiple subjects and invocations, such as an integration test. Managing state with the other kinds of test doubles is generally discouraged.
  - If they have a sensible API, offer a more natural way of asserting state. Rather than spying on specific calls to a dependency, you can query its final state to see if the real effect you want happened.
  - You can use them to run your application locally without spinning up or depending on real dependencies. This will usually improve developer experience (DX) because the fakes will be faster and more reliable than their real counterparts.

Spies, Mocks and Stubs can typically be autogenerated from an interface using a tool or using reflection. However, as Fakes encode the behavior of the dependency you're trying to make a double for, you'll have to write at least most of the implementation yourself

## [The problem with stubs and mocks](https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/working-without-mocks#the-problem-with-stubs-and-mocks)

> When you encode behavior into test doubles, you are adding your assumptions as to how the real dependency works into the test. If there is a discrepancy between the behavior of the double and the real dependency, or if one happens over time (e.g. the real dependency changes, which has to be expected), **you may have passing tests but failing software**.

Using fakes, **we can make assertions based on the final states of the respective systems rather than relying on complicated spying**. We'd ask each fake what records it held for the customer and assert they were updated. This feels more natural; if we manually checked our system, we would query those APIs to check their state, not inspect our request logs to see if we sent particular JSON payloads.

```go
// take our lego-bricks and assemble the system for the test
fakeAPI1 := fakes.NewAPI1()
fakeAPI2 := fakes.NewAPI2() // etc..
customerService := customer.NewService(fakeAPI1, fakeAPI2, etc...)

// create new customer
newCustomerRequest := NewCustomerReq{
	// ...
}
createdCustomer, err := customerService.New(newCustomerRequest)
assert.NoErr(t, err)

// we can verify all the details are as expected in the various fakes in a natural way, as if they're normal APIs
fakeAPI1Customer := fakeAPI1.Get(createdCustomer.FakeAPI1Details.ID)
assert.Equal(t, fakeAPI1Customer.SocialSecurityNumber, newCustomerRequest.SocialSecurityNumber)

// repeat for the other apis we care about

// update customer
updatedCustomerRequest := NewUpdateReq{SocialSecurityNumber: "123", InternalID: createdCustomer.InternalID}
assert.NoErr(t, customerService.Update(updatedCustomerRequest))

// again we can check the various fakes to see if the state ends up how we want it
updatedFakeAPICustomer := fakeAPI1.Get(createdCustomer.FakeAPI1Details.ID)
assert.Equal(t, updatedFakeAPICustomer.SocialSecurityNumber, updatedCustomerRequest.SocialSecurityNumber)
```

## Evolving software

Most software is not built and "finished" forever, in one release.

It's an incremental learning exercise, adapting to customer demands and other external changes. In the example, the APIs we were calling were also evolving and changing; plus, as we developed our software, we learned more about what system we really needed to make. Assumptions we made in our contracts turned out to be wrong or became wrong.

Thankfully, once the setup for the contracts was made, we had a simple way to deal with change. Once we learned something new, as a result of a bug being fixed or a colleague informing us that the API was changing, we'd:

- Write a test to exercise the new scenario. A part of this will involve changing the contract to drive you to simulate the behavior in the fake
- Running the test should fail, but before anything else, run the contract against the real dependency to ensure the change to the contract is valid.
- Update the fake so it conforms to the contract.
- Make the test pass.
- Refactor.
- Run all the tests and ship.

Running the full test suite before checking in may result in other tests failing due to the fake having a different behaviour. This is a good thing! You can now fix all the other areas of the system depending on the changed system; confident they will also handle this scenario in production. Without this approach, you'd have to remember to find all the relevant tests and update the stubs. Error-prone, laborious and boring.

# ent

- [An entity framework for Go](https://entgo.io/)

```sh
# Require `schema` folder and `generate.go`
cd adapters/driven/persistence/sqlite
go generate ./...
```
