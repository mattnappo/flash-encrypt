package crypto

import "testing"

func TestArgon2String(t *testing.T) {
	str := "test text"

	hash, err := Argon2String(str)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(hash)
}
