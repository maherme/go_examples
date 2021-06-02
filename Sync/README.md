# Sync
This code creates a counter which is safe to use concurrently.

## Mutex
Don't embed mutex in your struct since a mutex should not be modified outside your data structure:
```go
type Counter struct {
    sync.Mutex // don't do this, instead you should code it as: mu sync.Mutex
    value int
}

func (c *Counter) Inc() {
    c.Lock() // this should be c.mu.Lock()
    defer c.Unlock() // this should be defer c.mu.Unlock()
    c.value++
}
```
Reminder: a Mutex must not be copied after first use. You can look for this error using `vet` command.

## vet
You can use the `vet` command for doing some verification:
```console
go vet
```
