package crypto

import "testing"

func TestEncryptFile(t *testing.T) {
	passphrase := "secret passphrase"

	err := EncryptFile("./test/testfile", passphrase)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecryptFile(t *testing.T) {
	passphrase := "secret passphrase"

	err := DecryptFile("./test/testfile", passphrase)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEncryptDir(t *testing.T) {
	err := EncryptDir("./test", "secret passphrase")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecryptDir(t *testing.T) {
	err := DecryptDir("./test", "secret passphrase")
	if err != nil {
		t.Fatal(err)
	}
}