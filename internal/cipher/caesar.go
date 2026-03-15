package cipher

type Caesar struct {
	shift int
}

func NewCaesar(shift int) Cipher {
	shift = shift % 26
	if shift < 0 {
		shift += 26
	}
	return &Caesar{shift: shift}
}

func (c *Caesar) Encrypt(src []byte) []byte {
	dst := make([]byte, len(src))
	for i, b := range src {
		if b >= 'a' && b <= 'z' {
			dst[i] = 'a' + (b-'a'+byte(c.shift))%26
		} else if b >= 'A' && b <= 'Z' {
			dst[i] = 'A' + (b-'A'+byte(c.shift))%26
		} else {
			dst[i] = b
		}
	}
	return dst
}

func (c *Caesar) Decrypt(src []byte) []byte {
	inverseShift := 26 - c.shift
	decoder := &Caesar{shift: inverseShift}
	return decoder.Encrypt(src)
}
