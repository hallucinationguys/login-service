package token

import (
	"time"
)

type Maker interface {
	GenerateToken(username string, role string, duration time.Duration) (string, *Payload, error)
	ValidationToken(token string) (*Payload, error)
}
