package cipher

import (
	"bytes"
	"testing"
)

func TestOTP(t *testing.T) {
	tests := []struct {
		name string
		text string
		key  string
	}{
		{
			name: "Basic encryption",
			text: "HELLO WORLD",
			key:  "SUPERSECRET",
		},
		{
			name: "Japanese characters (multi-byte)",
			text: "こんにちは",
			key:  "secretkey",
		},
		{
			name: "Key shorter than text (modulo wrap test)",
			text: "This is a very long text",
			key:  "short",
		},
		{
			name: "Empty key (should return original text, not panic)",
			text: "Hello",
			key:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			textBytes := []byte(tt.text)
			keyBytes := []byte(tt.text)

			cipher := NewOTP(keyBytes)
			encrypted := cipher.Encrypt(textBytes)
			decrypted := cipher.Decrypt(encrypted)

			if !bytes.Equal(decrypted, textBytes) {
				t.Errorf("\nFailed: %s\nExpected: %q\nGot:    %q\n(Key: %q)",
					tt.name, tt.text, string(decrypted), tt.key)
			}
		})
	}
}
