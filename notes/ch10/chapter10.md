# Chapter 10: Error Handling

## Notes to Myself

+ Comparing errors directly doesn't work because they are pointers. But we
  also shouldn't directly compare the strings that two errors produce.
  Instead, we should wrap the errors and use `errors.Is` and errors.As` to
  compare what we got with what we want.  See [Comparing Go error
  values][comparing-go-error-values] and [Error wrapping in
  Go][error-wrapping-in-go] for more discussion.
+ Instead of constructing the error string the way he does on page 119,
  Siddiqui could use `fmt.Errorf` for cleaner code.
+ In addition, Siddiqui should not ignore possible errors by assigning them to
  the blank identifier, as he does on 119-120.  Instead, he should check for
  a non-`nil` error because if the error is not `nil` for a successful
  currency conversion, then something has gone wrong.

[comparing-go-error-values]: https://bitfieldconsulting.com/golang/comparing-errors
[error-wrapping-in-go]: https://bitfieldconsulting.com/golang/wrapping-errors
