package greetings

import (
	"fmt"
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