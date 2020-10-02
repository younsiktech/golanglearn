# golanglearn
learn golang

## INITIALIZATION for Developer

### Create Project

#### root folder always need those three folders
> mkdir bin && mkdir pkg && mkdir src

#### projects must create in src folder
> cd src && mkdir \<project name> && cd \<project name>

#### init module is required. (for installing go modules)
> go mod init \<github url>/\<project name>

### Module

#### for developer
> go get \<go module>

this command will modify go.mod file.

### TDD
> cd \<project_module>

#### Test all tests in module folder
> go test ./...

#### Test specify test in module folder
> go test ./\<module name>
>
>

## INITIALIZATION for user

### in Project
> go mod init

defined modules in go.mod will be installed