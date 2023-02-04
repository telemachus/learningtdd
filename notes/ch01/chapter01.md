# Chapter 1: The Money Problem

## Red, Green, Refactor

Siddiqui begins with an overview of the three-part technique of TDD.

1. *Red* Write a failing test and run it to prove that it fails.
1. *Green* Write just enough code to make the test pass and run the tests to
   prove that the code passes the previously failing test.
1. *Refactor* Clean up the code you've written.  If you break tests while
   refactoring, you must get things back to green before moving forward.

## The Money Problem

Siddiqui uses Kent Beck's money problem for the code examples.  We are writing
software that tracks money in multiple currencies.  The software should allow
users to do arithmetic on numbers in one currency and also to convert between
currencies.

## Note to Myself

Siddiqui uses the order `expected x, got y`, but I will follow Go's default of
`Function(val) = x, want y`.  See [Got before Want](go-test-wiki) on the Go
Test Wiki.

[go-test-wiki]: https://github.com/golang/go/wiki/TestComments#got-before-want
