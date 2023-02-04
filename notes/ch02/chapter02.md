# Chapter 2: Multicurrency Money

## Notes to Myself

On page 38, Siddiqui writes a test helper, but he doesn't call `t.Helper()`.
Initially, I added `t.Helper` to my code, but then I doubted myself.  Here's
what the [test wiki says](test-wiki-helpers):

> A test helper is a function that performs a setup or teardown task, such as
> constructing an input message, that does not depend on the code under test.
> ...
> Do not use this pattern when it obscures the connection between a test
> failure and the conditions that led to it.  Specifically, `t.Helper` should
> not be used to implement assert libraries.

Although Siddiqui is not creating an entire assert library, he is writing
(essentially) a specified assert function.  Nevertheless, that's not the key
issue.  If a test fails, I want the line number for the failure to reflect
where the helper is called (in the main test function) and not the test
helper. Thus, I do want `t.Helper`.

[test-wiki-helpers]: https://github.com/golang/go/wiki/TestComments#mark-test-helpers
