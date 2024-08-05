package main

import (
	"testing"
	"unicode/utf8"
)

// func TestReverse(t *testing.T) {
// 	testcases := []struct {
// 		in, want string
// 	}{
// 		{"Hello, world", "dlrow ,olleH"},
// 		{" ", " "},
// 		{"!12345", "54321!"},
// 	}
// 	for _, tc := range testcases {
// 		rev := Reverse(tc.in)
// 		if rev != tc.want {
// 			t.Errorf("Reverse: %q, want %q", rev, tc.want)
// 		}
// 	}
// }


/* The unit test has limitations, namely that each input must be added to the test by the developer. 
One benefit of fuzzing is that it comes up with inputs for your code, 
and may identify edge cases that the test cases you came up with didn’t reach.

we can keep unit tests, benchmarks, and fuzz tests in the same *_test.go file, 

When fuzzing, you can’t predict the expected output, since you don’t have control over the inputs.

two properties being checked in this fuzz test are:
1. Reversing a string twice preserves the original value
2. The reversed string preserves its state as valid UTF-8.

differences between the unit test and the fuzz test:
1. function begins with FuzzXxx instead of TestXxx, and takes *testing.F instead of *testing.T
2. Where we would expect to see a t.Run execution, 
we instead see f.Fuzz which takes a fuzz target function whose parameters are *testing.T and the types to be fuzzed. 
The inputs from your unit test are provided as seed corpus inputs using f.Add.

*/

func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
             return
        }
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
/* We can also run go test -run=FuzzReverse if We have other tests in that file, 
and we only wish to run the fuzz test.

Run FuzzReverse with fuzzing, to see if any randomly generated string inputs will cause a failure. 
This is executed using go test with a new flag, -fuzz, set to the parameter Fuzz.

Another useful flag is -fuzztime, which restricts the time fuzzing takes. 
e.g, specifying -fuzztime 10s in the test below would mean that, 
as long as no failures occurred earlier, the test would exit by default after 10 seconds had elapsed. 

to view the input that caused the failure, 
open the corpus file written to the testdata/fuzz/FuzzReverse directory in a text editor. 
Your seed corpus file may contain a different string, but the format will be the same.
The first line of the corpus file indicates the encoding version. 
Each following line represents the value of each type making up the corpus entry.

why fuzzreverse gave error:
The entire seed corpus used strings in which every character was a single byte. 
However, characters such as 泃 can require several bytes. 
Thus, reversing the string byte-by-byte will invalidate multi-byte characters.
*/