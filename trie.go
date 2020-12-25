package transliterate

// Test
type node struct {
	children map[rune]*node
	symbol   []rune
	phoneme  string
}

func newNode() *node {
	return &node{map[rune]*node{}, []rune{}, ""}
}

func (t *node) insert(symbol []rune, phoneme string) {
	for pos, char := range symbol {
		if t.children[char] != nil {
			t = t.children[char]
		} else {
			t.children[char] = newNode()
			t = t.children[char]
		}

		if pos == len(symbol)-1 {
			t.symbol = symbol
			t.phoneme = phoneme
			break
		}
	}
}
