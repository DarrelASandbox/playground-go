- [Question \& Answer](#question--answer)
  - [OS Exec](#os-exec)
  - [Error Types](#error-types)
  - [Context-Aware Reader](#context-aware-reader)

# Question & Answer

## OS Exec

keith6014 asks on [reddit](https://www.reddit.com/r/golang/comments/aaz8ji/testdata_and_function_setup_help/)

> I am executing a command using os/exec.Command() which generated XML data. The command will be executed in a function called GetData().
>
> In order to test GetData(), I have some testdata which I created.
>
> In my \_test.go I have a TestGetData which calls GetData() but that will use os.exec, instead I would like for it to use my testdata.
>
> What is a good way to achieve this? When calling GetData should I have a "test" flag mode so it will read a file ie GetData(mode string)?

- **Tips**:
  - When something is difficult to test, it's often due to the separation of concerns not being quite right
  - Don't add "test modes" into your code, instead use Dependency Injection so that you can model your dependencies and separate concerns.

## Error Types

Pedro on the Gopher Slack asks

> If I’m creating an error like `fmt.Errorf("%s must be foo, got %s", bar, baz)`, is there a way to test equality without comparing the string value?

- **Tips**:
  - If you find yourself testing for multiple error conditions don't fall in to the trap of comparing the error messages.
  - [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)

## Context-Aware Reader

This chapter demonstrates how to test-drive a context aware `io.Reader` as written by Mat Ryer and David Hernandez in [The Pace Dev Blog](https://pace.dev/blog/2020/02/03/context-aware-ioreader-for-golang-by-mat-ryer.html).

[In a previous chapter](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context) we discussed how we can use context to provide cancellation. This is especially useful if you're performing tasks which may be computationally expensive and you want to be able to stop them.

- What we want to demonstrate is something like
  - Given an `io.Reader` with "ABCDEF", when I send a cancel signal half-way through I when I try to continue to read I get nothing else so all I get is "ABC"
- So rather than reading everything, we could:
  - Supply a fixed-size byte array that doesn't fit all the contents
  - Send a cancel signal
  - Try and read again and this should return an error with 0 bytes read
- Summary:
  - Small interfaces are good and are easily composed
  - When you're trying to augment one thing (e.g io.Reader) with another you usually want to reach for the [delegation pattern](https://en.wikipedia.org/wiki/Delegation_pattern)

> In software engineering, the delegation pattern is an object-oriented design pattern that allows object composition to achieve the same code reuse as inheritance.

- An easy way to start this kind of work is to wrap your delegate and write a test that asserts it behaves how the delegate normally does before you start composing other parts to change behaviour. This will help you to keep things working correctly as you code toward your goal
