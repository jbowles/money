# Money
This is a rip-off of an old package called [finance](https://github.com/Confunctionist/finance) by `Confunctionist` on github. Credit is given in the `external` directory along with a copy of the original license.

Uses as its common value type `int64` but supports various operations for `float64`.


## Benchmarks

```sh
go test -bench Format

BenchmarkFormatUSD-8   	 1000000	      1944 ns/op
BenchmarkFormatI18USD-8	 1000000	      2314 ns/op
```
