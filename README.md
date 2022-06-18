## 💡 Overview
Go lang version of [GPUContainerRunner](https://github.com/KensukeNakazawa/GPUContainerRunner) for practicing go lang

## ✅ Setup

> git config commit.template .github/.commit_template

### :hammer: Build 
`local platform`
> go build -o rungpu cmd/rungpu/main.go  

`linux platform`
> GOOS=linux GOARCH=amd64 go build -o rungpu cmd/rungpu/main.go
