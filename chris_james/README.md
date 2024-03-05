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
```

It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. Every test has a cost.

Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer

Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)

Adding methods so you can add functionality to your data types and so you can implement interfaces

[Table driven tests](https://go.dev/wiki/TableDrivenTests) to make your assertions clearer and your test suites easier to extend & maintain

At some point you may wish to use structs to manage state, exposing methods to let users change the state in a way that you can control.

Pointers to struct have their own name: struct pointers and they are [automatically dereferenced](https://go.dev/ref/spec#Method_values)
