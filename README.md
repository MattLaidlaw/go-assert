# Go-Assert
A collection of familiar test assertions for use with Go's built-in testing environment.

## Installation

## Supported Assertions
The following assertions cause the testing object 't' to fail if evaluated to false
```go
// assert that the first two parameters are identical objects
ExpectEq(expected interface{}, actual interface{}, t *testing.T)
```
