package slices

import (
	"slices"
)

type Runes []rune

func (r Runes) Read(offset int, limit int) Runes {
	if offset < 0 || limit < 0 {
		return nil
	}
	if offset+limit > len(r) {
		return nil
	}

	return slices.Clone(r[offset : offset+limit])
}

func (r Runes) ReadString(offset int, limit int) string {
	return r.Read(offset, limit).String()
}

func (r Runes) Index(sub []rune) int {
	return Index(r, sub)
}

func (r Runes) FindString(s string) int {
	return r.Index([]rune(s))
}

func (r Runes) StringArray() []string {
	var result []string
	for i := range r {
		result = append(result, string(r[i]))
	}
	return result
}

func (r Runes) String() string {
	return string(r)
}

func StringToRunes(s string) Runes {
	return Runes(s)
}
