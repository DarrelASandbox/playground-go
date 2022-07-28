<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#basics">Basics</a>
      <ol>
        <li><a href="#grouping-data">Grouping Data</a></li>
      </ol>
    </li>
  </ol>
</details>

&nbsp;

## About The Project

- Learn How To Code: Google's Go (golang) Programming Language
- The Ultimate Comprehensive Course - Perfect for Both Beginners and Experienced Developers

&nbsp;

---

&nbsp;

## Basics

- [The Go Blog - Using Go Modules](https://go.dev/blog/using-go-modules)

```sh
# Creating a new module
go mod init example.com/username/repo

# Adding a dependency
go get rsc.io/quote

# Upgrading dependencies
go list -m all
go get golang.org/x/text
go get rsc.io/sampler
```

- [The Go Blog - Get familiar with workspaces](https://go.dev/blog/get-familiar-with-workspaces)

> <b>Daniel: </b>Go has two organization tools: modules and packages.
>
> <b>A package is a group of Go files in the same directory.</b> All Go files in the same directory must have the same package statement at the top of the file (one exception in tests), and they can all access all variables, constants, types, and functions from each other. Go files in one package must use an import statement to access values from other packages.
>
> <b>A module is a group of packages.</b> It is defined by a go.mod file, and all go packages in the current directory and sub-directories are a part of that module, unless they have their own go.mod file. A module has a name that looks like an import path. To import a package in another module, you must have the other module listed in your module (dependency management). When you use an import statement to get a package, you include the name of the module in the statement.
>
> For example, I have a module named github.com/myname/awesomemodule and it has a sub-folder named stringtools. To import stringtools, I use import "github.com/myname/awesomemodule/stringtools" If the code attempting to import stringtools is not a part of awesomemodule, I need to add awesomemodule to my go.mod, probably by using go get github.com/myname/awesomemodule from the command line while in one of the directories in my module.

&nbsp;

---

&nbsp;

- Control flow
  - sequence
  - loop; iterative
  - conditional
- [Short variable declarations - declare and assign](https://go.dev/ref/spec#Short_variable_declarations)
- [Types - composite types](https://go.dev/ref/spec#Type_declarations)
- [Conversions](https://go.dev/ref/spec#Conversions)

> <b>Oren: </b>type hotdog int
>
> <b>Todd: </b>The main idea here is that VALUES are of a certain TYPE. You can have the value "TODD" which is of type STRING. You can have the value 42 which is of type INT. When we create our own types, we can specify the UNDERLYING type. This allows us to create code that is self-documenting, among other things
>
> So when you see <code>type hotdog int</code> or <code>type person struct</code>, we have create the TYPE hotdog and person. We can then create VALUES of TYPE hotdog or person. So in our program we might have <code>var x hotdog = 42</code> which would tell us we have 42 hotdogs.
>
> THE MAIN USE of creating our own types, however, is not for something trivial like "hotdog". It is for creating SEMANTICS in our code and for creating data structures, like type person, for aggregating (also known as, gathering, composing) data together. :)

&nbsp;

---

&nbsp;

- [Go Spec - Iota](https://go.dev/ref/spec#Iota)
- [Wiki - Extended Backus–Naur form](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form)

&nbsp;

### Grouping Data

- <b>array</b>
  - a numbered sequence of elements of a single type
  - does not change in size
  - used for Go internals; generally not recommended for your code
- <b>slice</b>
  - built on top of an array
  - changes in size
  - holds values of the same type
  - has a length and a capacity
  - Slices are built on top of arrays. A slice is dynamic in that it will grow in size. The underlying array, however, does not grow in size. When we create a slice, we can use the built in function make to specify how large our slice should be and also how large the underlying array should be. This can enhance performance a little bit.
- <b>map</b>
  - key/value storage
  - an unordered group of elements of one type, called the element type, indexed
    by a set of unique keys of another type, called the key type.
- <b>struct</b>
  - a data structure
  - a composite type
  - allows us to collect values of different types together

&nbsp;

- [Go Standard Library](https://pkg.go.dev/std)
- [Go Forum - Are arrays passed by value or passed by reference in GO?](https://forum.golangbridge.org/t/are-arrays-passed-by-value-or-passed-by-reference-in-go/19255)

&nbsp;

---

&nbsp;

> <b>Matteo: </b>I don't understand what you mean with Overriding: "promotions"
>
> <b>Jared: </b>In a previous lesson, he mentioned that "promotion" means you can access the fields of the "inner" type directly, which equates to them being promoted. For example, in his example he had created a type "person" which he then used as an inner type in another type called "secret agent". To access the fields in "person", he didn't have to go through "person", like "s1.person.first". He could access the fields directly, such as "s1.first". I know it's been a year since you asked the question, but it helps me to remember better if I type out an answer :)

&nbsp;

---

&nbsp;

- <code>func (r receiver) identifier(parameters) (return(s)) {...}</code>
- we define our func with parameters (if any)
- we call our func and pass in arguments (in any)
- <b>Everything in Go is PASS BY VALUE</b>
- A <b>variadic parameter</b> is a func which takes an unlimited number of arguments.
- A <b>method</b> is nothing more than a FUNC attached to a TYPE. When you attach a func to a
  type it is a method of that type. You attach a func to a type with a RECEIVER.
- <b>Interfaces & polymorphism</b>
  - In Go, values can be of more than one type.
  - An interface allows a value to be of more than one type. We create an interface using this syntax: “keyword identifier type” so for an interface it would be: “type human interface”
  - We then define which method(s) a type must have to implement that interface. If a TYPE has the required methods, which could be none (the empty interface denoted by interface{}), then that TYPE implicitly implements the interface and is also of that interface type.
  - In Go, values can be of more than one type.

&nbsp;

> <b>Dharlequin: </b>What is the actual idea/concept behind using defer?

> <b>Daniel: </b>Defer is useful for two circumstances. First is making sure a function is run before the surroundings exit, no matter which return statement is used to do the exiting. Go conventions prefers designing your conditions to exit a function early, instead of other languages (mostly c/c++) method of only using a single return at the end of the function. Which makes defer nice since you don't need to worry about the cleanup at every return.
>
> Defer is also makes sure it runs when a panic occurs. So if something in your function ends up panicking, but it is then recovered further up the call stack, the defer will still have run to get the cleanup done. In this way, you can consider defer similar to other languages concept of try/finally.

&nbsp;

---

&nbsp;

> <b>Abhishek Kumar: </b>what is the purpose of anonymous function? In what kind of scenario can I use them?

> <b>Todd: </b>You might want to "encapsulate" a variable; narrow the scope of some variables; so you could put it all in a function. You also might want to "abstract" out some code; get it modularized so that you could say "this piece does X, this piece does Y, this piece does Z" where each piece was an anonymous func. You also might want to launch a goroutine so you could do this: <code>go func(){<code here>}()</code> Some of the uses!

&nbsp;

---

&nbsp;

- <b>Callback</b>
  - passing a func as an argument
  - functional programming not something that is recommended in go, however, it is good to be aware of callbacks
  - idiomatic go: write clear, simple, readable code
- <b>Closure</b>
  - passing a func as an argument
  - [functional programming](https://stackoverflow.com/questions/66839360/does-go-support-functional-programming) not something that is recommended in go, however, it is good to be aware of callbacks
  - idiomatic go: write clear, simple, readable code

&nbsp;

---

&nbsp;
