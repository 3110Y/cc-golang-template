//go:build wireinject
// +build wireinject

package di

//go:generate wire

import (
	"github.com/3110Y/profile/internal/infrastructure/database"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

type DI struct {
	DB *sqlx.DB
}

func NewDI(
	DB *sqlx.DB,
) *DI {
	return &DI{
		DB: DB,
	}
}

func InitializeDI() (*DI, error) {
	wire.Build(
		NewDI,
		database.NewConnect,
	)
	return &DI{}, nil
}
