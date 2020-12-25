package transliterate

// Converter is
type Converter struct {
	trie *node
}

// NewConverter creates new converter
func NewConverter(key map[string]string) (c *Converter) {
	c = &Converter{}
	c.AssignKey(key)
	return
}

// AssignKey is
func (c *Converter) AssignKey(key map[string]string) {
	c.trie = newNode()
	for symbol, phoneme := range key {
		c.trie.insert([]rune(symbol), phoneme)
	}
}

// Convert is
func (c Converter) Convert(text string) (result string) {
	if c.trie == nil {
		return text
	}

	characters := []rune(text)
	root := c.trie

	for pos := 0; pos < len(characters); pos++ {
		t := root
		depth := -1
		phoneme := ""

		for end := pos; end < len(characters); end++ {
			char := characters[end]
			if t.children[char] == nil {
				break
			}
			if t.children[char].symbol != nil {
				depth = end - pos
				phoneme = t.children[char].phoneme
			}
			t = t.children[char]
		}

		if depth >= 0 {
			result += phoneme
			pos += depth
		} else {
			result += string(characters[pos])
		}
	}

	return
}
