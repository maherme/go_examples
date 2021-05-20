# Hello World

## Quick start
To run the application:
```console
go run hello.go
```
To run the test:
```console
go test
```
If you have not got the go.mod file you need to do:
```console
go mod init hello
```
## Test file
Rules for writing a test:
- It needs to be in a file with a name like xxx_test.go.
- The test function must start with the word Test.
- The test function takes one argument only t *testing.T.
- In order to use the *testing.T type, you need to import "testing".

t.Helper() is needed to tell the test suite that this method is a helper. So when a test fails the line number reported will be in our function call rather than inside our test helper. 
