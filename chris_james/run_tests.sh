#!/bin/bash

# chmod +x run_tests.sh

directories=(
    "./01-go-fundamentals/01-hello-world"
    "./01-go-fundamentals/02-integers"
    "./01-go-fundamentals/03-iteration"
    "./01-go-fundamentals/04-arrays-and-slices"
    "./01-go-fundamentals/05-structs-methods-interfaces"
    "./01-go-fundamentals/06-pointers-and-errors"
    "./01-go-fundamentals/07-maps"
    "./01-go-fundamentals/08-dependency-injection"
    "./01-go-fundamentals/09-mocking"
    "./01-go-fundamentals/10-concurrency"
    "./01-go-fundamentals/11-racer"
    "./01-go-fundamentals/12-reflection"
    "./01-go-fundamentals/13-sync"
    "./01-go-fundamentals/14-context"
    "./01-go-fundamentals/15-intro-to-property-based-tests"
    "./01-go-fundamentals/16-clockface"
    "./01-go-fundamentals/17-blogposts"
    "./01-go-fundamentals/18-templating"
    "./01-go-fundamentals/19-generics"
    "./01-go-fundamentals/20-generics-and-arrays"
)

for dir in "${directories[@]}"; do
    echo "Running tests in $dir"
    (cd "$dir" && go test ./...)
done