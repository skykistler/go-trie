package trie

import "testing"

func TestBasicWords(t *testing.T) {
	testTrie := NewTrie()

	words := [...]string{
		"test",
		"aardvark",
		"alpha",
		"beta",
		"gamma",
	}

	for _, word := range words {
		testTrie.Insert(word)
	}

	testWords := map[string]bool{
		"":         false,
		"alpha":    true,
		"zeta":     false,
		"aardvark": true,
		"test":     true,
		"sigma":    false,
		"beta":     true,
		"theta":    false,
		"gamma":    true,
	}

	for word, expected := range testWords {
		if contains := testTrie.Contains(word); contains != expected {
			t.Errorf("Got %t when expected %t for value: %s", contains, expected, word)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	testTrie := NewTrie()

	words := [...]string{
		"",
		"!@#$%",
		"twice",
		"twice",
	}

	for _, word := range words {
		testTrie.Insert(word)
	}

	testWords := map[string]bool{
		"":      true,
		"!":     false,
		"!@#$%": true,
		"twice": true,
	}

	for word, expected := range testWords {
		if contains := testTrie.Contains(word); contains != expected {
			t.Errorf("Got %t when expected %t for value: %s", contains, expected, word)
		}
	}

}
