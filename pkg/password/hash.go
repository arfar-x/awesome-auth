package password

import "golang.org/x/crypto/bcrypt"

func Check(value1 string, value2 string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(value1), []byte(value2)); err != nil {
		return false
	}
	return true
}

func Make(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), 4)
	if err != nil {
		panic("Password hash generation failed.")
	}
	return string(hash), nil
}
