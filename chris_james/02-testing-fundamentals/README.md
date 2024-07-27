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

# shell

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

# First system: HTTP API

1. A **driver**. In this case, one works with an HTTP system by using an **HTTP client**. This code will know how to work with our API. Drivers translate DSLs into system-specific calls; in our case, the driver will implement the interface specifications define.
2. An **HTTP server** with a greet API
3. A **test**, which is responsible for managing the life-cycle of spinning up the server and then plugging the driver into the specification to run it as a test

## Adapter Pattern

> In software engineering, the adapter pattern is a software design pattern (also known as wrapper, an alternative naming shared with the decorator pattern) that allows the interface of an existing class to be used as another interface.[1] It is often used to make existing classes work with others without modifying their source code.

## Reflect

- Analyze your problem and identify a slight improvement to your system that pushes you in the right direction
- Capture the new essential complexity in a specification
- Follow the compilation errors until the AT runs
- Update your implementation to make the system behave according to the specification
- Refactor

## Complexity

In the context of Go testing, **essential complexity** refers to the inherent challenges and intricacies that arise directly from the problem you're trying to solve. This includes the logic of the tests, the conditions you need to cover, and the actual behavior of the system under test. For example, if you're testing a complex algorithm or a system with many interdependent components, the complexity of ensuring all paths are tested correctly is essential.

**Accidental complexity**, on the other hand, arises from the tools, frameworks, or practices used to implement the tests, rather than from the problem itself. In Go, this could include difficulties with setting up the test environment, managing dependencies, or dealing with limitations or quirks in the testing framework. For instance, if your tests are hard to write or maintain because of the way the test setup is structured, that’s accidental complexity.

The goal in Go testing, as with any testing, is to minimize accidental complexity so that you can focus on the essential complexity—ensuring your code works correctly under all necessary conditions. Go's standard testing package is designed to keep accidental complexity low by providing simple and effective tools for writing tests.

## High-level steps for acceptance test

- Build a docker image
- Wait for it to be listening on some port
- Create a driver that understands how to translate the DSL into system specific calls
- Plug in the driver into the specification

## New functionality

- Shouldn't have to change the specification;
- Should be able to reuse the specification;
- Should be able to reuse the domain code.

# Second System: GRPC

```sh
# Generate the client and server code
cd adapters/grpcserver
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    greet.proto
```

## Overview

- `adapters` have cohesive units of functionality grouped together
- `cmd` holds our applications and corresponding acceptance tests
- Our code is totally decoupled from any accidental complexity

## When should I write acceptance tests?

1. Is this an edge case? I'd prefer to unit test those
2. Is this something that the non-computer people talk about a lot? I would prefer to have a lot of confidence the key thing "really" works, so I'd add an acceptance test
3. Am I describing a user journey, rather than a specific function? Acceptance test
4. Would unit tests give me enough confidence? Sometimes you're taking an existing journey that already has an acceptance test, but you're adding other functionality to deal with different scenarios due to different inputs. In this case, adding another acceptance test adds a cost but brings little value, so I'd prefer some unit tests.

# Third System: Web

- [GopherCon UK 2021: Riya Dattani & Chris James - Acceptance Tests, BDD & GO](https://www.youtube.com/watch?v=ZMWJCk_0WrY)

> Imagine the least technical person that you can think of, who understands the problem-domain, reading your Acceptance Tests. The tests should make sense to that person.
