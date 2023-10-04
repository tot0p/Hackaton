package hash

import "testing"

func TestHashPassword(t *testing.T) {

	s := "password"

	hashed, err := HashPassword(s)
	if err != nil {
		t.Error(err)
	}

	if len(hashed) == 0 {
		t.Error("hashed password is empty")
	}
}

func TestCheckPassword(t *testing.T) {

	s := "password"

	hashed, err := HashPassword(s)
	if err != nil {
		t.Error(err)
	}

	if len(hashed) == 0 {
		t.Error("hashed password is empty")
	}

	ok := CheckPasswordHash(s, hashed)
	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Error("passwords do not match")
	}
}
