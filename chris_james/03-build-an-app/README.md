- [Intro](#intro)

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

# Command Line & Project Structure

# Time

# WebSockets
