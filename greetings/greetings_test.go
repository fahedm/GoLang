package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.

// Test function names have the form TestName, where Name says something about the specific test.
// test functions take a pointer to the testing package's testing.T type as a parameter. 
// You can use this parameter's methods for reporting and logging from your test.
func TestHelloName(t *testing.T) {
	/*
	TestHelloName calls the Hello function, 
	passing a name value with which the function should be able to return a valid response message. 
	If the call returns an error or an unexpected response message (one that doesn't include the name you passed in), 
	you use the t parameter's Fatalf method to print a message to the console and end execution.
	*/
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello_with_Exception("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	/*
	TestHelloEmpty calls the Hello function with an empty string. 
	This test is designed to confirm that your error handling works. 
	If the call returns a non-empty string or no error, 
	you use the t parameter's Fatalf method to print a message to the console and end execution.
	*/
    msg, err := Hello_with_Exception("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

// "go test" command will execute test functions (names begins with Test) in test files (names end with _test)
// Ending a file's name with _test.go tells the go test command that this file contains test functions.

// "go test -v" -v flag to get verbose output that lists all of the tests and their results.