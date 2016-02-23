[![Go Report Card](https://goreportcard.com/badge/github.com/jbowles/money)](https://goreportcard.com/report/github.com/jbowles/money)
[![GoDoc](https://godoc.org/github.com/jbowles/money?status.svg)](https://godoc.org/github.com/jbowles/money)

# Money


This is a rip-off of an old package called [finance](https://github.com/Confunctionist/finance) by `Confunctionist` on github. Credit is given in the `external` directory along with a copy of the original license.

Uses as its common value type `int64` but supports various operations for `float64`.


## Todo
Check overflow for multiplication and division.


## Testing
Trying to get thoroughly tested. Also, I always forget some of the test commands:

```sh
go test -ovf=0 ## custom test flag to run overflow test
go test
go test -v
go test -v -run Mul

go test -cover
go tool cover -func=coverage.out
go test -covermode=count -coverprofile=count.out fmt
go tool cover -func=count.out
go tool cover -html=coverage.out
```


## Benchmarks
I alwasy forget how to run bemchmarks!

```sh
go test -bench=.
go test -bench Format  #or some regular expression
```
