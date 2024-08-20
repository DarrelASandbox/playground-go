- [Intro](#intro)
- [HTTP Server](#http-server)
- [JSON, Routing and Embedding](#json-routing-and-embedding)
  - [Why not test the JSON string?](#why-not-test-the-json-string)
- [IO and Sorting](#io-and-sorting)
  - [Didn't we just break some rules there? Testing private things? No interfaces?](#didnt-we-just-break-some-rules-there-testing-private-things-no-interfaces)
- [Command Line \& Project Structure](#command-line--project-structure)
- [Time](#time)
- [WebSockets](#websockets)

# Intro

- **HTTP server** - We will create an application which listens to HTTP requests and responds to them.
- **JSON, routing and embedding** - We will make our endpoints return JSON and explore how to do routing.
- **IO and sorting** - We will persist and read our data from disk and we'll cover sorting data.
- **Command line & project structure** - Support multiple applications from one code base and read input from command line.
- **Time** - using the time package to schedule activities.
- **WebSockets** - learn how to write and test a server that uses WebSockets.

# HTTP Server

```sh
go build && ./http-server
curl -X POST http://localhost:5000/players/Pepper
curl -X POST http://localhost:5000/players/Pepper
curl -X POST http://localhost:5000/players/Pepper
curl http://localhost:5000/players/Pepper
```

By adding mutexes, we enforce concurrency safety especially for the counter in our RecordWin function.

# JSON, Routing and Embedding

## Why not test the JSON string?

- **Brittleness**. If you change the data-model your tests will fail.
- **Hard to debug**. It can be tricky to understand what the actual problem is when comparing two JSON strings.
- **Poor intention**. Whilst the output should be JSON, what's really important is exactly what the data is, rather than how it's encoded.
- **Re-testing the standard library**. There is no need to test how the standard library outputs JSON, it is already tested. Don't test other people's code.

# IO and Sorting

## [Didn't we just break some rules there? Testing private things? No interfaces?](https://quii.gitbook.io/learn-go-with-tests/build-an-application/io#didnt-we-just-break-some-rules-there-testing-private-things-no-interfaces)

It's true that in general you should favour not testing private things as that can sometimes lead to your tests being too tightly coupled to the implementation, which can hinder refactoring in future.

However, we must not forget that tests should give us confidence.

We were not confident that our implementation would work if we added any kind of edit or delete functionality. We did not want to leave the code like that, especially if this was being worked on by more than one person who may not be aware of the shortcomings of our initial approach.

Finally, it's just one test! If we decide to change the way it works it won't be a disaster to just delete the test but we have at the very least captured the requirement for future maintainers.

# Command Line & Project Structure

```sh
go test
cd cmd/webserver && go run main.go
curl -X POST http://localhost:5000/players/Pepper
curl -X POST http://localhost:5000/players/Pepper
curl -X POST http://localhost:5000/players/Pepper
curl http://localhost:5000/players/Pepper
# Check http://localhost:5000/league
```

# Time

- [Dummy objects are passed around but never actually used. Usually they are just used to fill parameter lists.](https://martinfowler.com/articles/mocksArentStubs.html)

Remember that any type can implement an interface, not just `structs`. If you are making a library that exposes an interface with one function defined it is a common idiom to also expose a `MyInterfaceFunc` type.

This type will be a `func` which will also implement your interface. That way users of your interface have the option to implement your interface with just a function; rather than having to create an empty `struct` type.

# WebSockets
