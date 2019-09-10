# Go Modules Tutorial

In go programing, it was very tough to maintain the versioning of go package. Let's suppose that we have an application and we are using some package with v1.0 version. Later package version upgraded to v2.0 with lot of changes. This will affect our application. So we need something that maintains package versions.
Go 1.11 adds preliminary support for a new concept called “modules”. It supports versioning and package distribution.
go command enables the use of modules when the current directory or any parent directory has a go.mod.


  - Creating a new module.
  - Adding a dependency
  - Upgrading dependencies
  - Removing unused dependencies

## Creating a new module

To create a module, just run the following commands

```
 $ go mod init <module name/path>
```
OR
```
$ go mod init
```

If above commands give warnings " modules disabled by GO111MODULE=auto", then run the following commands
```
 $ GO111MODULE=on go mod init <module name/path>
```
OR
```
$ GO111MODULE=on go mod init
```
 go mod init command wrote a go.mod file. Only go.mod file will be there at root folder of the module.
 
 ### Adding a dependency
go build, go test, and other package-building commands add new dependencies to go.mod as needed.

### Upgrading dependencies
In the code example, you can see that I'm using untaged version of github.com/chilts/sid by go list -m all.
Let's upgrade this
```
$ go get github.com/chilts/sid
```
This will get latest version of github.com/chilts/sid.

### Removing unused dependencies
Run the following command to remove unused dependencies
```
$ go mod tidy
```