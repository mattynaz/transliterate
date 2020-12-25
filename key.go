package transliterate

import (
	"encoding/csv"
	"fmt"
	"os"
)

// // IpaSimplifierKey is
// var IpaSimplifierKey = map[string]string{
// 	".":    ".",
// 	"a":    "a",
// 	"aɪ":   "ai",
// 	"aʊ":   "ow",
// 	"aː":   "a",
// 	"b":    "b",
// 	"d":    "d",
// 	"dz":   "dz",
// 	"dʒ":   "j",
// 	"e":    "ey",
// 	"eː":   "ay",
// 	"f":    "f",
// 	"h":    "h",
// 	"i":    "ee",
// 	"iː":   "ee",
// 	"i̯":   "ee",
// 	"j":    "y",
// 	"k":    "k",
// 	"l":    "l",
// 	"l̩":   "l",
// 	"m":    "m",
// 	"m̩":   "m",
// 	"n":    "n",
// 	"n̩":   "n",
// 	"o":    "o",
// 	"oː":   "o",
// 	"p":    "p",
// 	"pf":   "pf",
// 	"r":    "r",
// 	"s":    "s",
// 	"t":    "t",
// 	"ts":   "ts",
// 	"tɕ":   "ch",
// 	"tʃ":   "ch",
// 	"u":    "oo",
// 	"uː":   "oo",
// 	"u̯":   "w",
// 	"v":    "v",
// 	"w":    "w",
// 	"x":    "kh",
// 	"y":    "u",
// 	"yː":   "ew",
// 	"z":    "z",
// 	"ã":    "ã",
// 	"ãː":   "aw",
// 	"æ":    "ah",
// 	"ç":    "hy",
// 	"ð":    "th",
// 	"õ":    "õ",
// 	"õː":   "o",
// 	"ø":    "u",
// 	"øː":   "ew",
// 	"øːr":  "er",
// 	"ŋ":    "ŋ",
// 	"œ":    "uh",
// 	"œːɐ̯": "er",
// 	"œ̃":   "uh",
// 	"œ̃ː":  "uh",
// 	"ɐ":    "u",
// 	"ɐ̯":   "ɐ̯",
// 	"ɑ":    "aw",
// 	"ɑ̃":   "ãw",
// 	"ɔ":    "o",
// 	"ɔʊ̯":  "ow",
// 	"ɔʏ":   "oy",
// 	"ɔ̃":   "aw",
// 	"ɕː":   "sh",
// 	"ə":    "uh",
// 	"ɛ":    "e",
// 	"ɛɪ̯":  "ay",
// 	"ɛː":   "ai",
// 	"ɛ̃":   "ũh",
// 	"ɛ̃ː":  "eh",
// 	"ɟʝ":   "j",
// 	"ɡ":    "g",
// 	"ɣ":    "ɣ",
// 	"ɥ":    "w",
// 	"ɨ":    "e",
// 	"ɪ":    "i",
// 	"ɫ":    "w",
// 	"ɱ":    "m",
// 	"ɲ":    "ny",
// 	"ɵ":    "oo",
// 	"ɹ":    "r",
// 	"ɾ":    "r",
// 	"ʁ":    "r",
// 	"ʂ":    "sh",
// 	"ʃ":    "sh",
// 	"ʉ":    "oo",
// 	"ʊ":    "w",
// 	"ʎ":    "yl",
// 	"ʏ":    "oo",
// 	"ʐ":    "zh",
// 	"ʑː":   "zh",
// 	"ʒ":    "zh",
// 	"ʔ":    "'",
// 	"ʝ":    "y",
// 	"ʲ":    "ʸ",
// 	"ˈ":    "ˈ",
// 	"ˌ":    "ˌ",
// 	"ː":    "ː",
// 	"͡":    "͡",
// 	"β":    "b",
// 	"θ":    "th",
// }

// LoadKey loads
func LoadKey(path string) (key map[string]string, err error) {
	csvfile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer csvfile.Close()

	lines, err := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	key = map[string]string{}
	if len(lines) == 0 || len(lines[0]) != 2 {
		return nil, fmt.Errorf("invalid key format, file must be CSV with two columns and one row for the column headers")
	}
	for _, line := range lines[1:] {
		fmt.Println(line[0] + ", " + line[1])
		key[line[0]] = line[1]
	}

	return
}
