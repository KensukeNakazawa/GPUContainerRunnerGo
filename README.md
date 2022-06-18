## 💡 Overview
Go lang version of [GPUContainerRunner](https://github.com/KensukeNakazawa/GPUContainerRunner) for practicing go lang

## ✅ Setup

> git config commit.template .github/.commit_template

### :hammer: Build 
> go build cmd/rungpu/main.go && mv main rungpu
> GOOS=linux GOARCH=amd64 go build -o rungpu cmd/rungpu/main.go