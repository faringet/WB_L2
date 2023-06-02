package main

import "testing"

func TestUnpackS1(t *testing.T) {
	input := "a4bc2d5e"
	expOut := "aaaabccddddde"
	funcOut := UnpackS(input)

	if funcOut != expOut {
		t.Errorf("UnpackS() = %q, expOut %q", funcOut, expOut)
	}
}

func TestUnpackS2(t *testing.T) {
	input := "abcd"
	expOut := "abcd"
	funcOut := UnpackS(input)

	if funcOut != expOut {
		t.Errorf("UnpackS() = %q, expOut %q", funcOut, expOut)
	}
}

func TestUnpackS3(t *testing.T) {
	input := "45"
	expOut := ""
	funcOut := UnpackS(input)

	if funcOut != expOut {
		t.Errorf("UnpackS() = %q, expOut %q", funcOut, expOut)
	}
}

func TestUnpackS4(t *testing.T) {
	input := ""
	expOut := ""
	funcOut := UnpackS(input)

	if funcOut != expOut {
		t.Errorf("UnpackS() = %q, expOut %q", funcOut, expOut)
	}
}
