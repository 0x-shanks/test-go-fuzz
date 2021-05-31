package main

import (
	"testing"
)

func FuzzFibo(f *testing.F) {
	f.Add(2)

	f.Fuzz(func(t *testing.T, n int) {
		if Fibo(n) != Fibo(n-2)+Fibo(n-1) {
			t.Fatal("NG")
		}
	})
}
