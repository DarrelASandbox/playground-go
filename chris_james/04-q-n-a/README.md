- [Question \& Answer](#question--answer)
  - [OS Exec](#os-exec)
  - [Error Types](#error-types)

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
