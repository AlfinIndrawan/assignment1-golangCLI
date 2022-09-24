package helpers

import "golang.org/x/crypto/bcrypt"

// menghasilkan hash dari password sebelum disimpan ke database
func HashPassword(password string) string {
	salt := 8
	pass := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(pass, salt)
	return string(hash)
}

// mengkomparasi password yang diinputkan dengan password yang sudah dihash
func ComparePass(h, p string) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
