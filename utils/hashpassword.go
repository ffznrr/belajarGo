package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil {
		return "", err
	}
	
	return string(hashedByte),err
}

func CheckPasswordHash(password, hashPassword string)bool{
	res := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return res == nil
}