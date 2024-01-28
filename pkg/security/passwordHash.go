package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	bytePassword := []byte(password)
	byteHashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)

	return string(byteHashedPassword), err

}

func DoPasswordsMatch(hashedPassword, currentPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currentPassword))
	return err == nil

}
