- [shell](#shell)
- [Introduction To Acceptance Tests](#introduction-to-acceptance-tests)
  - [What are they?](#what-are-they)
  - [Benefits of acceptance tests](#benefits-of-acceptance-tests)
  - [Potential drawbacks vs unit tests](#potential-drawbacks-vs-unit-tests)
  - [High-level steps for the acceptance test](#high-level-steps-for-the-acceptance-test)
  - [`LaunchTestProgram`](#launchtestprogram)
- [Scaling Acceptance Tests](#scaling-acceptance-tests)
  - [Anatomy of bad acceptance tests](#anatomy-of-bad-acceptance-tests)
  - [Tight coupling](#tight-coupling)
- [specifications](#specifications)
  - [First system: HTTP API](#first-system-http-api)

# shell

```sh
# Run all tests, ignoring cached results
go test -count=1 ./...
```

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

# specifications

## First system: HTTP API

1. A **driver**. In this case, one works with an HTTP system by using an **HTTP client**. This code will know how to work with our API. Drivers translate DSLs into system-specific calls; in our case, the driver will implement the interface specifications define.
2. An **HTTP server** with a greet API
3. A **test**, which is responsible for managing the life-cycle of spinning up the server and then plugging the driver into the specification to run it as a test

```sh
go test -v ./cmd/httpserver

# Test the Container Manually
docker build -t greeter-server -f cmd/httpserver/Dockerfile .
docker run -p 8080:8080 greeter-server
curl http://localhost:8080/greet
```
