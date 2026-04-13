package main

import "testing"

func TestStrLen(t *testing.T) {
	got := strLen("hello world")
	want := 11

	if got != want {
		t.Errorf("strLen(\"hello world\") = %d; want %d", got, want)
	}
}
