package greetings

import (
	"errors"
	"fmt"
)

//Hello returns a greeting for the named person

/* A function whose name starts with capital letter can be called by a function not in the same package.
It is called Exported name
func <FunctionName>(<ParameterName> <ParameterType>) <Function ReturnType>
*/

func Hello(name string) string{ 
	//return a greeting that embeds the name in message
	message := fmt.Sprintf("Hi, %v. Let's Learn Go", name)
	// := is the operator for declaring and intializing a variable in one line
	return message
}

func Hello_with_Exception(name string) (string, error){
	// If no name is given, return a error with message
	if name == "" {
		return "", errors.New("empty name")
	}
	// If name is given, return a value that embeds the name in greeting message
	message := fmt.Sprintf("Hi, %v, Let's Learn Go", name)
	return message, nil
}