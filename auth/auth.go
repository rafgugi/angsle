package auth

import (
	"github.com/rafgugi/angsle/entity"
	"github.com/rafgugi/angsle/sec"
)

var (
	publicKey  = "pub"
	privateKey = "s3cr3t"
)

type userRepo interface {
	ByCredential(username, password string) entity.User
}

// Auth TODO
type Auth struct {
	userRepo  userRepo
	tokenRepo tokenRepo
}

type tokenRepo interface {
	TokenByOwner(id uint64) string
	CipherByToken(token string) string
	Store(token string, ownerID uint64, cipher string)
}

// GenerateToken TODO
func (a *Auth) GenerateToken(username, password string) string {
	user := a.userRepo.ByCredential(username, password)
	token := a.tokenRepo.TokenByOwner(user.ID)
	if token != "" {
		return token
	}

	bearer := sec.EncryptUser(user, privateKey, publicKey)
	token = sec.Sha256(bearer)
	a.tokenRepo.Store(token, user.ID, bearer)
	return token
}

// ByToken TODO
func (a *Auth) ByToken(token string) *entity.User {
	bearer := a.tokenRepo.CipherByToken(token)
	if bearer == "" {
		return nil
	}

	user := sec.DecryptUser(bearer, publicKey)
	return &user
}

// RefreshToken TODO
func (a *Auth) RefreshToken(token string)
