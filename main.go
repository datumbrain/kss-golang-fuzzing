package main

func equal(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func divide(a int, b int) int {
	if b == 0 {
		return -1
	}
	return a / b
}
