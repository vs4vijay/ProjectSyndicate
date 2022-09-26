# ProjectSyndicate


---

## Tech. Stack

- https://github.com/zeromicro/go-zero

---

### Development Guidelines

```go

go mod init syndicate

go get -u github.com/tal-tech/go-zero

go install github.com/tal-tech/go-zero/tools/goctl@latest

// Generate *.api file
goctl api -o healthcheck.api
goctl rpc new greet

// Generate API from *.api
goctl api go -api healthcheck.api -dir healthcheck

```