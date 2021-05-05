Package exit implements errors represents the status code.
======================================================================

Purpose
----------------------------------------------------------------------

1. eliminating a boiler plate.
2. specifying an status code.


1 eliminating a boiler plate.
----------------------------------------------------------------------

The package exit eliminates a boiler plate for main() like below.

### without using the package exit

``` go
func main() {
  if err := run(); err != nil {
    fmt.Fprintf(os.Stderr, "error: %v\n", err)
    os.Exit(1)
  }
}
```

### with using the package exit

``` go
func main() {
  exit.ExitOnError(run())
}
```

2 specifying an status code.
----------------------------------------------------------------------

### the status code WITHOUT the error message

Using `exit.Status`, `exit.Exit` does not write any error messages, calls just os.Exit(42).

``` go
exit.Exit(exit.Status(42))
```

### the status code WITH the error message

Using `exit.Error`, `exit.Exit` writes the error message of err, then calls os.Exit(42).

``` go
exit.Exit(exit.Error(42, errors.New("Deep Thought said")))
```
