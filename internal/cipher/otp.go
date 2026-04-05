package cipher

type OTP struct {
	key []byte
}

func pad(src []byte, key []byte) []byte {
	if len(src) == 0 || len(key) == 0 {
		return src
	}
	result := make([]byte, len(src))
	for i := range src {
		result[i] = src[i] ^ key[i%len(key)]
	}
	return result
}

func NewOTP(key []byte) Cipher {
	return &OTP{key: key}
}

func (o *OTP) Encrypt(src []byte) []byte {
	return pad(src, o.key)
}

func (o *OTP) Decrypt(src []byte) []byte {
	return pad(src, o.key)
}
