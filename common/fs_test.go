package common

import "testing"

func TestListDir(t *testing.T) {
	files, err := ListDir("../")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(files)
}