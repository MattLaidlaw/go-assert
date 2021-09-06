# Go-Assert
A collection of familiar test assertions for use with Go's built-in testing environment.

## Installation
```
go get github.com/MattLaidlaw/go-assert@d96695b3290cfd16bba4cccf2163f3894341e84b
```

## Supported Assertions
The following assertions cause the testing object 't' to fail if evaluated to false
```go
// assert that the first two parameters are identical objects
ExpectEq(actual interface{}, expected interface{}, t *testing.T)
```
