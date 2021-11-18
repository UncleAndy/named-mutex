# named-mutex

Named mutex for go

Sample:

```go
    mutex := NamedMutex("any name of mutex")
    mutex.Lock()
    defer mutex.Unlock()
```
