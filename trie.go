package trie

// Trie represents a node (or root node) of a prefix trie
type Trie struct {
	children map[string]Trie
	terminal bool
}

// NewTrie will return a new root node Trie
func NewTrie() Trie {
	return Trie{
		children: make(map[string]Trie),
		terminal: false,
	}
}

// Insert will create trie nodes for each character in the value string
func (t *Trie) Insert(value string) bool {
	if len(value) == 0 {
		return t.setTerminal()
	}

	first := string(value[0])

	child, contains := t.children[first]
	if !contains {
		child = NewTrie()
		t.children[first] = child
	}

	if len(value) > 1 {
		return child.Insert(value[1:])
	} else {
		changed := child.setTerminal()
		t.children[first] = child
		return changed
	}
}

// Contains will return true or false if value has an exact match in the Trie
func (t *Trie) Contains(value string) bool {
	if len(value) == 0 {
		return t.terminal
	}

	key := string(value[0])

	if child, contains := t.children[key]; !contains {
		return false
	} else {
		if len(value) == 1 {
			return child.terminal
		} else {
			return child.Contains(value[1:])
		}
	}
}

func (t *Trie) setTerminal() bool {
	changed := !t.terminal
	if changed {
		t.terminal = true
	}
	return changed
}
