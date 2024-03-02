- [quii](https://github.com/quii)

```sh
go test
godoc -http=:6060 # http://localhost:6060/pkg/

# Benchmarking
# Benchmarks are run sequentially.
go test -bench=.

go test -cover
```

It is important to question the value of your tests. It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. Having too many tests can turn in to a real problem and it just adds more overhead in maintenance. Every test has a cost.

- [Table Driven Tests](https://go.dev/wiki/TableDrivenTests)
