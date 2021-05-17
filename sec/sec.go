package sec

import "github.com/rafgugi/angsle/entity"

// EncryptUser returns cipher, encryption of user
func EncryptUser(user entity.User, privateKey, publicKey string) string

// DecryptUser returns user, from given cipher text
func DecryptUser(cipher, publickKey string) entity.User

func Sha256(msg string) string
