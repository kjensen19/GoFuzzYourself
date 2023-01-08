package main

import (
	"testing"
	"unicode/utf8"
)

// Fuzz run with $ go test -fuzz=Fuzz will run until an error or when stopped wigth ctrl - c
// $ go test -fuzz=Fuzz -fuzztime 30s will run for 30s and then stop

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) //use f.Add to provide "seed corpus" (see terms dictionary in SWE bookmarks)
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
			t.Errorf("Before; %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produce invalid UTF-8 string %q", rev)
		}
	})
}

//Basic Unit test, pre-fuzz conversion
// func TestReverse(t *testing.T) {
// 	testcases := []struct {
// 		in, want string
// 	}{
// 		{"hello, world", "dlrow ,olleh"},
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
