package pwd

import "golang.org/x/crypto/bcrypt"

/*
Hash hashes the password
*/
func Hash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

/*
Verify verifies the password
*/
func Verify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
