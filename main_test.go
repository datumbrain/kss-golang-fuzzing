package main

import "testing"

func FuzzEqual(f *testing.F) {
	f.Add([]byte{1, 2, 3}, []byte{1, 2, 3})
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		equal(a, b)
	})
}

func FuzzDivide(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int, b int) {
		divide(a, b)
	})
}
