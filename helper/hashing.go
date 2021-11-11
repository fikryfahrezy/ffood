package helper

import "golang.org/x/crypto/bcrypt"

func CompareHash(hash, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))

	return err == nil
}

func GenerateHash(text string) (string, bool) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)

	return string(hashed), err == nil
}
