# Arrays and Slices
## Arrays
You can create an array in two ways:
```go
numbers := [5]int{1, 2, 3, 4, 5}
numbers := [...]int{1, 2, 3, 4, 5}
```
## Slices
They are similar to arrays, but you do not specify the size:
```go
array := []int{1, 2, 3, 4, 5}
```
You can create slices using the built-in function make:
```go
slice := make([]int, 4) // This will create [0 0 0 0]
```
## Test coverage
You can know the coverage of your test using:
```console
go test -cover
```
More info: https://blog.golang.org/slices-intro
