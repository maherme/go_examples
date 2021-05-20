# Iteration

## For loop
This is a code example with the for loop syntax:
```go
package main

import "fmt"

func main() {

    // The most basic type, with a single condition.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
    
    // A classic initial/condition/after for loop.
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }
    
    // for without a condition will loop repeatedly until you break out 
    // of the loop or return from the enclosing function.
    for {
        fmt.Println("loop")
        break
    }

    // You can also continue to the next iteration of the loop.
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
```
## Benchmark
You can watch an example of benchmark in [repeat_test.go](Iteration/repeat_test.go).

You do not need to specify the number of times the code is run; the framework determines that for a good result.

For executing the benchmark:
```console
go test -bench=.
```
My result using this code:
```console
[maherme@localhost Iteration]$ go test -bench=.
goos: linux
goarch: amd64
pkg: iteration
cpu: Intel(R) Core(TM) i7-4700HQ CPU @ 2.40GHz
BenchmarkRepeat-8   	 6690025	       198.8 ns/op
PASS
ok  	iteration	1.516s
```
