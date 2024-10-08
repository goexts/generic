package slices

import (
	"slices"
)

type Bytes []byte

func (r Bytes) Read(offset int, limit int) Bytes {
	if offset < 0 || limit < 0 {
		return nil
	}
	if offset+limit > len(r) {
		return nil
	}

	return slices.Clone(r[offset : offset+limit])
}

func (r Bytes) ReadString(offset int, limit int) string {
	return r.Read(offset, limit).String()
}

func (r Bytes) Index(sub []byte) int {
	return Index(r, sub)
}

func (r Bytes) FindString(s string) int {
	return r.Index([]byte(s))
}

func (r Bytes) StringArray() []string {
	result := make([]string, 0, len(r))
	for i := range r {
		result = append(result, string(r[i]))
	}
	return result
}

func (r Bytes) String() string {
	return string(r)
}

func StringToBytes(s string) Bytes {
	return Bytes(s)
}
