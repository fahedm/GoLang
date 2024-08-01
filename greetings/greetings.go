package greetings

import (
	"fmt"
	"errors"
)

//Hello returns a greeting for the named person

/* A function whose name starts with capital letter can be called by a function not in the same package.
It is called Exported name
func <FunctionName>(<ParameterName> <ParameterType>) <Function ReturnType>
*/

func Hello(name string) string {
	// Return a greeting with name
	message := fmt.Sprintf("Hi, %v. Welcomee", name)
	// := is the operator for declaring and intializing a variable in one line
	// Using fmt package's Sprintf function to create a greeting message. 
	// The first argument is a format string, and Sprintf substitutes the name parameter's value for the %v format verb. 
	// Inserting the value of the name parameter completes the greeting text.
	return message
}

// go mod init command creates a go.mod file to track your code's dependencies

func Hello_with_Exception(name string) (string, error){
	// If no name is given, return a error with message
	if name == "" {
		return "", errors.New("empty name")
	}
	// If name is given, return a value that embeds the name in greeting message
	message := fmt.Sprintf("Hi, %v, Let's Learn Go", name)
	return message, nil
	// nil (meaning no error) as a second value in the successful return
}