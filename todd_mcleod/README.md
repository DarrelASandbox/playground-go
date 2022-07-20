<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
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
