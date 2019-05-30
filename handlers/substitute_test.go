package handlers

import (
	"testing"
)

func setupTokens() (tokens map[string]string) {
	tokens = make(map[string]string)
	tokens["a"] = "a"
	tokens["b"] = "b"
	tokens["c"] = "c"
	return tokens
}

func TestSubstituteStart(t *testing.T) {
	tokens := setupTokens()
	e := "abc"
	a := Substitute("${a}bc", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteMiddle(t *testing.T) {
	tokens := setupTokens()
	e := "abc"
	a := Substitute("a${b}c", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteEnd(t *testing.T) {
	tokens := setupTokens()
	e := "abc"
	a := Substitute("ab${c}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteAll(t *testing.T) {
	tokens := setupTokens()
	e := "abc"
	a := Substitute("${a}${b}${c}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteEscape(t *testing.T) {
	tokens := setupTokens()
	e := "a$${b}c"
	a := Substitute("${a}$${b}${c}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}
