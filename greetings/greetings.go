package greetings

import (
	"fmt"
	"errors"
	"math/rand" // to generate a random number for selecting an item from the slice.
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

func Hello_with_random(name string) (string, error){
	if name == "" {
		return "", errors.New("empty name")
	}
	// create message using random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.

// randomFormat starts with a lowercase letter, 
// making it accessible only to code in its own package (in other words, it's not exported).
func randomFormat() string{
	// A slice of message formats.

	/* In randomFormat, declare a formats slice with three message formats. 
	When declaring a slice, you omit its size in the brackets, like this: []string. 
	This tells Go that the size of the array underlying the slice can be dynamically changed. */
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]

}