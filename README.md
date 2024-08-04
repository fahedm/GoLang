# GoLang

Learning GoLang

## Branches

- main : initial commit
- develop : Initial Learning
- fahed : basics
- docs-tutorial : Learning from GoLang Documentation Tutorial

## Compile and Install

- to compile the program we use "go build" command
- in hello directory, we will run this command
- after compiling, "./hello" command from hello directory will execute your program("hello.exe" on windows)
- to avoid the execution from executable directory or with its path, we need to install the program.
- to install the program, first we need to find install path, that will be found by "go list"
- go list -f '{{.Target}}' > this will give you the install path
- now, add this path to your system's shell path. > $ export PATH=$PATH:/path/to/your/install/directory
- after PATH is updated, run "go install" to compile and install the package
- now try to run the name of the program > hello
- If it is not executing we need to check for PATH in the system, something is not correctly set.
- Below are some commands to fix the Path
  
  - export GOPATH=$HOME/go
  - export GOROOT=/usr/local/go
  - export GOBIN=$GOPATH/bin
  - export PATH=$PATH:$GOPATH
  - export PATH=$PATH:$GOROOT/bin
  - export PATH=$PATH:$GOPATH/bin

- after these commands run "go install" again, now you should be able to execute you program name as command

## Web Service Implementation

- we will use Gin framework for this.
- as of now, we will use in-memory database and simple curl command to test our functionality.
- a real-world service would likely use a database query to perform this lookup.
  
## Generics

- With generics, you can declare and use functions or types that are written to work with any of a set of types provided by calling code.
- 