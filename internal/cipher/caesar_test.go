package cipher

import (
	"testing"
)

// test shifts more than 26
// for example: shift(29) == shift(3)
func TestShiftMoreThan26(t *testing.T) {
	plaintext := []byte("encryption")

	cipher29 := NewCaesar(29)
	cipher3 := NewCaesar(3)

	res29 := string(cipher29.Encrypt(plaintext))
	res3 := string(cipher3.Encrypt(plaintext))

	if res29 != res3 {
		t.Errorf("Expected shift 29 to equal shift 3. Got '%s' and '%s'", res29, res3)
	}
}

// plaintext == shift(0) == shift(26)
func TestShift26(t *testing.T) {
	plaintext := []byte("Hello")

	cipher26 := NewCaesar(26)
	cipher0 := NewCaesar(0)

	res26 := string(cipher26.Encrypt(plaintext))
	res0 := string(cipher0.Encrypt(plaintext))

	if res26 != res0 {
		t.Errorf("Expected shift 26 to equal shift 0. Got '%s' and '%s'", res26, res0)
	}
	if res26 != "Hello" {
		t.Errorf("Expected shift 26 to equal plaintext 'Hello'. Got '%s'", res26)
	}
}

func TestUpperLowerCase(t *testing.T) {
	plaintext := []byte("HelloWorld")
	shift := 5
	expected := "MjqqtBtwqi"

	c := NewCaesar(shift)
	cipher := c.Encrypt(plaintext)

	if string(cipher) != expected {
		t.Errorf("Encryption failed: expected '%s', got '%s'", expected, string(cipher))
	}

	decrypted := c.Decrypt(cipher)
	if string(decrypted) != string(plaintext) {
		t.Errorf("Decryption failed: expected '%s', got '%s'", string(plaintext), string(decrypted))
	}
}

// test chars ^[a-zA-Z] in plaintext
// ignores out-of-range chars and puts them to output directly
func TestCharsOutRange(t *testing.T) {
	plaintext := "12 3 4 @ ^% ^^[]"
	shift := 122

	c := NewCaesar(shift)
	cipher := c.Encrypt([]byte(plaintext))

	if string(cipher) != plaintext {
		t.Errorf("Encryption failed: expected '%s', got '%s'", plaintext, string(cipher))
	}

	decrypted := string(c.Decrypt(cipher))
	if decrypted != plaintext {
		t.Errorf("Decryption failed: expected '%s', got '%s'", plaintext, decrypted)
	}
}

// test shift((-x) % 26) == shift((26-x) % 26)
func TestNegativeMod(t *testing.T) {
	x := 1234
	plaintext := []byte("HelloWorld")

	shiftNeg := (-x) % 26
	shiftNonNeg := (26 - x) % 26

	cNeg := NewCaesar(shiftNeg)
	cNonNeg := NewCaesar(shiftNonNeg)

	cipherNeg := string(cNeg.Encrypt(plaintext))
	cipherNonNeg := string(cNonNeg.Encrypt(plaintext))

	if cipherNeg != cipherNonNeg {
		t.Errorf("Expected negative mod logic to match. Got '%s' and '%s'", cipherNeg, cipherNonNeg)
	}
}

func TestNegativeShift(t *testing.T) {
	plaintext := []byte("HelloWorld")
	shift := -3
	expected := "EbiilTloia"

	c := NewCaesar(shift)
	cipher := c.Encrypt(plaintext)

	if string(cipher) != expected {
		t.Errorf("Encryption failed: expected '%s', got '%s'", expected, string(cipher))
	}

	decrypted := string(c.Decrypt(cipher))
	if decrypted != string(plaintext) {
		t.Errorf("Decryption failed: expected '%s', got '%s'", string(plaintext), decrypted)
	}
}
