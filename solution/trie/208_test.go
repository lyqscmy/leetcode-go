package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	root := Constructor()
	root.Insert("abc")
	root.Insert("abd")
	if root.Search("a") {
		t.Error("should be false")
	}
	if !root.Search("abc") {
		t.Error("should be true")
	}
	if root.StartsWith("ac") {
		t.Error("should be false")
	}
	if !root.StartsWith("a") {
		t.Error("should be true")
	}
}
