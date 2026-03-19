// simple (unkeyed) columnar transposition
// both decrypt and encrypt methods use runes to handle non-english characters

package cipher

type UnkeyedColumnarTransposition struct {
	column int
}

func (u *UnkeyedColumnarTransposition) Decrypt(src []byte) []byte {
	if u.column == 1 || len(src) == 0 {
		return src
	}

	runes := []rune(string(src))
	dstRunes := make([]rune, len(runes))

	srcIdx := 0
	for c := range u.column {
		for i := c; i < len(runes); i += u.column {
			dstRunes[i] = runes[srcIdx]
			srcIdx++
		}
	}
	return []byte(string(dstRunes))
}

func (u *UnkeyedColumnarTransposition) Encrypt(src []byte) []byte {
	if u.column == 1 || len(src) == 0 {
		return src
	}

	runes := []rune(string(src))
	dstRunes := make([]rune, 0, len(runes))

	for c := range u.column {
		for i := c; i < len(runes); i += u.column {
			dstRunes = append(dstRunes, runes[i])
		}
	}

	return []byte(string(dstRunes))
}

func NewUnkeyedColumnarTransposition(column int) Cipher {
	if column <= 1 {
		column = 1
	}
	return &UnkeyedColumnarTransposition{column: column}
}
