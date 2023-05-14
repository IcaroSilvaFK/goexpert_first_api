package utils

import "golang.org/x/crypto/bcrypt"

func VerifyHash(plain, hash string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))

	return err
}
