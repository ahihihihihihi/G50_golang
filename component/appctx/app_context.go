package appctx

import "gorm.io/gorm"

type Appcontext interface {
	GetMainDBConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB  {
	return ctx.db
}
