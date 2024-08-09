package infrastructure

import "golang.org/x/crypto/bcrypt"

func HashePassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return " ", err
	}

	return string(hashedPassword), nil

}

func ComparePassword(existingPassword string, newPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(newPassword))

	if err != nil {
		return false
	}

	return true

}
