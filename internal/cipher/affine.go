package cipher

type Affine struct {
	a int
	b int
}

func NewAffine(a int, b int) *Affine {
	return &Affine{a: a, b: b}
}

func (c *Affine) Encrypt(src []byte) []byte {
	dst := make([]byte, len(src))
	for i, b := range src {
		if b >= 'a' && b <= 'z' {
			val := int(b - 'a')
			dst[i] = byte((c.a*val+c.b)%26) + 'a'
		} else if b >= 'A' && b <= 'Z' {
			val := int(b - 'A')
			dst[i] = byte((c.a*val+c.b)%26) + 'A'
		} else {
			dst[i] = b
		}
	}
	return dst
}

func (c *Affine) Decrypt(src []byte) []byte {
	dst := make([]byte, len(src))

	aInv := modInverse(c.a, 26)

	for i, b := range src {
		if b >= 'a' && b <= 'z' {
			val := int(b - 'a')
			res := (aInv * (val - c.b)) % 26
			if res < 0 {
				res += 26
			}
			dst[i] = byte(res) + 'a'
		} else if b >= 'A' && b <= 'Z' {
			val := int(b - 'A')
			res := (aInv * (val - c.b)) % 26
			if res < 0 {
				res += 26
			}
			dst[i] = byte(res) + 'A'
		} else {
			dst[i] = b
		}
	}
	return dst
}

func modInverse(a, m int) int {
	a = a % m
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return 1
}