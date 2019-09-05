package crypto

import "testing"

func TestEncryptFile(t *testing.T) {
	passphrase := "secret passphrase"

	err := EncryptFile("./testfile", passphrase)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecryptFile(t *testing.T) {
	passphrase := "secret passphrase"

	err := DecryptFile("./testfile", passphrase)
	if err != nil {
		t.Fatal(err)
	}
}