package main // main package to group functions, package is made up of all files in same directory

import (
	"fmt" // fmt is very popular package, which deal with texts. Standard Library Package
	// "rsc.io/quote" // an external package, published in modules where other packages can use them
	"example.com/greetings" // another package created in root directory.
)

// main function will execute when we run the main package

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println("Let's learn GoLang from Documentation")
	// fmt.Println(quote.Go())

	message := greetings.Hello("Fahed")
	fmt.Println(message)
}

// to initialize a package : go mod init <package name>
// to run the first Go Program : go run .
// to get list of other commands : go help 

// to add new modeule requirements and sums : go mod tidy

// When you ran go mod tidy, 
// it located and downloaded the rsc.io/quote module that contains the package you imported. 
// By default, it downloaded the latest version

// go mod edit -replace <declaredPackage>=../<packageName>
// This specifies that example.com/greetings should be replaced with ../greetings 
// for the purpose of locating the dependency.

// go mod tidy after every change in imports 
// The command found the local code in the greetings directory, 
// then added a require directive to specify that example.com/hello requires example.com/greetings. 
// You created this dependency when you imported the greetings package in hello.go.