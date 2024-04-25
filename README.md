# Hash and Verify Password

## Installation

```bash
go get -u github.com/metadiv-io/pwd
```

## Hash Password

```go
hash := pwd.Hash("password")
```

## Verify Password

```go
ok := pwd.Verify("password", hash)
```
