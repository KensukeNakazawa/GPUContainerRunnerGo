## 💡 Overview
Go lang version of [GPUContainerRunner](https://github.com/KensukeNakazawa/GPUContainerRunner) for practicing go lang

version: `go1.18.3 darwin/arm64`

## :wrench: Setup

> $ git config commit.template .github/.commit_template

### :hammer: Build 
`local platform`
> $ go build -o rungpu cmd/rungpu/main.go  

`linux platform`
> $ GOOS=linux GOARCH=amd64 go build -o rungpu cmd/rungpu/main.go

### :white_check_mark: Test
`all`
> $ go test -v ./...

`under package`
> ex)  
> $ go test -v ./pkg/gpu