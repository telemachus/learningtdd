# TODO items

+ Improve how the library implements errors in conversion rates (e.g., create
  a custom error type that is a slice of strings; store the specific
  conversion failures there; implement an `Error` method that returns
  a reasonable string representation of those failures)
+ Improve the implementation of exchange rates
+ Allow exchange rates to be modified
