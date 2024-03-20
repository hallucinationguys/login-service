package components

import (
	"github.com/hallucinationguys/login-service/pkg/components/token"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	GetTokenMaker() token.Maker
}

type appCtx struct {
	db         *gorm.DB
	secret     string
	tokenMaker token.Maker
}

func NewAppContext(db *gorm.DB, secretKey string, tokenMaker token.Maker) *appCtx {
	return &appCtx{db: db, secret: secretKey, tokenMaker: tokenMaker}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db.Debug()
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secret
}

func (ctx *appCtx) GetTokenMaker() token.Maker {
	return ctx.tokenMaker
}
