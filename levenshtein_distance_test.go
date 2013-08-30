package main

import (
	"testing"
)

func TestCalculatesDistance(t *testing.T) {
	l := NewLevenshteinDistance()
	l.AddWord("hello")
	l.AddWord("there")
	l.AddWord("good")
	l.AddWord("sirs")

	result := l.FindCloseWords("hell", 3)
	if len(result) != 2 {
		t.Fatalf("Expected 2 results")
	}

	expectItem(t, result[0], "hello", 1)
	expectItem(t, result[1], "there", 3)
}

func expectItem(t *testing.T, item LevenshteinWordMatch, word string, distance int) {
	if item.Word != word {
		t.Fatalf("Expected word to be %q but was %q", word, item.Word)
	}
	if item.Distance != distance {
		t.Fatalf("Expected distance to be %v but was %v", distance, item.Distance)
	}
}
