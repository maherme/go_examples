# Mocking
This application prints a countdown with a configurable timeout between the digit printed. It will print Go! when countdown finishes.
```console
[maherme@localhost Mocking]$ go run countdown.go
3
2
1
Go![maherme@localhost Mocking]$ 
```
You should read the [package time](https://golang.org/pkg/time/) for a better understanding of this module. Specially type [duration](https://golang.org/pkg/time/#Duration) and func [sleep](https://golang.org/pkg/time/#Sleep).

In this example we are mocking ```time.Sleep```, so we can use dependency injection to use it instead of a "real" ```time.Sleep``` and then we can spy on the calls to make assertions on them.

We are defining our dependency as an interface. This lets us then use a real Sleeper in main and a spy sleeper in our tests.
