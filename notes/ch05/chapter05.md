# Chapter 5: Packages and Modules in Go

## Notes to Myself

On page 67, Siddiqui implies that if you import a package as `tdd/stocks`, you
would need to prefix `tdd/stocks` every time that you use code from that
package.  In fact, you only have to use the last part of an import.  In this
case, you would use, e.g., `stocks.Money` not `tdd/stocks.Money`.
