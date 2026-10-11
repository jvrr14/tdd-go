# Learn Go with Tests

My exercises from [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/).
Each chapter has its own folder. The code and its tests stay together:

```text
go.mod                      Shared Go settings for this whole repository
main.go                     Runnable example for the Hello chapter
hello/
  hello.go                  Greeting code
  hello_test.go             Tests for the greeting code
integers/
  adder.go                  Addition code
  adder_test.go             Tests for addition
iteration/
  iteration.go              Repetition code
  iteration_test.go         Tests and a benchmark for repetition
```

## Run the exercises

Run these commands from the repository folder (the folder containing `go.mod`):

```sh
go test ./hello              # Test the Hello chapter
go test ./integers           # Test the Integers chapter
go test ./iteration          # Test the Iteration chapter
go test ./...                # Test all chapters
```

For more detail about each test, add `-v`, for example `go test -v ./hello`.

## Run the Hello example

From the repository folder, run:

```sh
go run .
```

The dot means "this folder"; Go requires a file or package after `go run`.
This prints `Hello, World!`. Open `main.go` and change the arguments to
`hello.Hello("Javokhir", "Spanish")` to try a different name and language.

## What does `package` mean?

Think of a package as a group of Go files that work together. In this repository,
each chapter folder is one group. The first line of each Go file names that group:

- Files in `hello/` start with `package hello`.
- Files in `integers/` start with `package integers`.
- Files in `iteration/` start with `package iteration`.

The code and test files use the same package name, so a test can call a function
such as `Hello(...)` or `Add(...)` directly.

`package main` is Go's special name for a program you launch yourself. Its
`func main()` is where that program starts. In `main.go`, it calls the
Hello chapter's greeting function and prints the result. The import line tells
Go where to find that chapter's code; `hello.Hello(...)` calls its `Hello` function.

The chapter folders contain the functions you're learning and their tests.
`go test` handles starting those tests for you. The small program in `main.go`
lets you also run the greeting yourself, just like the book's Hello example.

## Follow the Arrays and Slices chapter

1. Create an `arrays/` folder next to the other chapter folders.
2. Follow the book by writing `arrays/sum_test.go` first.
3. Use `package arrays` as its first line, where the book uses `package main`.
4. Run `go test ./arrays` and read the failure.
5. When the book asks you to add `sum.go`, create `arrays/sum.go` with
   `package arrays` as its first line too.
6. Continue the book's cycle: write a test, see it fail, write enough code to pass,
   then tidy up while keeping the tests passing.

Keep using the existing `go.mod` for all chapters. You don't need to run
`go mod init` again inside a chapter folder.

For later function exercises, follow the same pattern: one folder per chapter,
with matching package names in its code and test files. When a chapter builds an
actual runnable program, follow its instructions for `package main` and `main()`.

## Reading the tests

The tests use the book's naming convention:

- `got` is the result returned by your code.
- `want` is the result the test expects.

Files ending in `_test.go` contain tests. An `Example...` function with an
`Output:` comment is also checked by `go test`. A `Benchmark...` function measures
speed; run the existing benchmark with `go test -bench=. ./iteration`.
