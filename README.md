# OldDriver

### Spider, Make life more enjoyable.

### Project Init
```shell
go mod tidy
go mod vendor
```

### Mac 下编译

```shell
# linux 下去执行
CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build main.go
# Windows 下去执行
CGO_ENABLED=0 GOOS=windows  GOARCH=amd64  go  build  main.go
```
### Linux 下编译
```shell
# Mac  下去执行
CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64  go build main.go
# Windows 下执行
CGO_ENABLED=0 GOOS=windows  GOARCH=amd64  go build main.go
```
### Windows 下编译
```shell
# Mac 下执行
SET  CGO_ENABLED=0
SET  GOOS=darwin
SET  GOARCH=amd64
go   build main.go

# Linux 去执行
SET  CGO_ENABLED=0
SET  GOOS=linux
SET  GOARCH=amd64
go   build main.go
```