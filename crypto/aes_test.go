package crypto

import "testing"

func TestGenerateKey(t *testing.T) {
	key, err := GenerateKey("secret passphrase")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(*key)
}

func TestEncryptFile(t *testing.T) {
	key, err := GenerateKey("secret passphrase")
	if err != nil {
		t.Fatal(err)
	}

	err = EncryptFile("./test/testfile", key)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecryptFile(t *testing.T) {
	key, err := GenerateKey("secret passphrase")
	if err != nil {
		t.Fatal(err)
	}

	err = DecryptFile("./test/testfile", key)
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
