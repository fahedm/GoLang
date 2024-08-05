package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	// rev := Reverse(input)
	// doubleRev := Reverse(rev)
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	// fmt.Printf("reversed: %q\n", rev)
	// fmt.Printf("reversed again: %q\n", doubleRev)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)

}

// function will accept a string, loop over it a byte at a time, and return the reversed string at the end.
func Reverse(s string) (string, error) {
	// b := []byte(s)
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	fmt.Printf("input: %q\n", s)
	b := []rune(s) // traverse the string by runes, instead of by bytes.
	fmt.Printf("input: %q\n", b)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	// This code is based on the stringutil.Reverse function within golang.org/x/example.
	// return string(b)
	return string(s), nil
}
