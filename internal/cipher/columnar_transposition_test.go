package cipher

import (
	"testing"
)

type UnkeyedColumnarTranspositionTest struct {
	name      string
	plaintext string
	column    int
	expected  string
}

func TestUnkeyedColumnarTransposition(t *testing.T) {
	tests := []UnkeyedColumnarTranspositionTest{
		{
			name:      "Even grid length",
			plaintext: "secret",
			column:    3,
			expected:  "sreect",
		},
		{
			name:      "Uneven grid length",
			plaintext: "secrets",
			column:    3,
			expected:  "srseect",
		},
		{
			name:      "Multi-byte characters (rune)",
			plaintext: "Şerif Göktuğ", // turkish and space
			column:    3,
			expected:  "ŞiGteföur kğ",
		},
		{
			name:      "Column larger than plaintext",
			plaintext: "secret",
			column:    7,
			expected:  "secret",
		},
		{
			name:      "Single column fallback",
			plaintext: "hello",
			column:    1,
			expected:  "hello",
		},
		{
			name:      "Negative column fallback",
			plaintext: "hello",
			column:    -1,
			expected:  "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cipher := NewUnkeyedColumnarTransposition(tt.column)

			encrypted := cipher.Encrypt([]byte(tt.plaintext))
			if string(encrypted) != tt.expected {
				t.Errorf("Encrypt() = %v, want %v", string(encrypted), tt.expected)
			}

			decrypted := cipher.Decrypt(encrypted)
			if string(decrypted) != tt.plaintext {
				t.Errorf("Decrypt() = %v, want %v", string(decrypted), tt.plaintext)
			}
		})
	}
}
