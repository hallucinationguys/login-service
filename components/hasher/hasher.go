package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

type bcryptHash struct{}

func NewbcryptHash() *bcryptHash {
	return &bcryptHash{}
}

func (h *bcryptHash) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	return string(hashedPassword), nil
}

func (h *bcryptHash) VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
