# Chapter 11: Banking on Redesign

## Notes to Myself

The `assertNil` test helper should also use `t.Helper` at the top of the
function.  See my notes on `assertEqual` in Chapter 2.

The `Convert` function on page 135 would be more idiomatic if we replace `if
ok { ... }` with `if !ok { ... }`. That way, the happy path stays on the left.
