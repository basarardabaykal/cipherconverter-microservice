package cipher

type Cipher interface {
	Encrypt(src []byte) []byte
	Decrypt(src []byte) []byte
}
