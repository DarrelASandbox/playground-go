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

# IO and Sorting

# Command Line & Project Structure

# Time

# WebSockets
