package cypher

import (
	"strings"
)

const alphabet = "\"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZабвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ\""

type Cryptographer struct {
	alphabetMap   map[rune]int
	alphabetRunes []rune
}

func New() *Cryptographer {
	m := make(map[rune]int)
	for i, r := range []rune(alphabet) {
		m[r] = i
	}
	return &Cryptographer{m, []rune(alphabet)}
}

func (c *Cryptographer) Encrypt(s string, k int) string {
	sb := strings.Builder{}
	for _, r := range s {
		runePosition, exist := c.alphabetMap[r]
		if exist {
			p := (runePosition + k) % len(c.alphabetRunes)
			sb.WriteRune(c.alphabetRunes[p])
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func (c *Cryptographer) Decrypt(s string, k int) string {

	sb := strings.Builder{}

	for _, r := range s {
		runePosition, exist := c.alphabetMap[r]
		if exist {
			if runePosition-k < 0 {
				runePosition = len(c.alphabetRunes) - runePosition - k
			}
			sb.WriteRune(c.alphabetRunes[runePosition-k])
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
