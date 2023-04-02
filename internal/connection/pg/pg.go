package pg

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

type PGConfig struct {
	ConnectionString     string `default:"postgres://postgres:postgres@localhost:5432/promo_db"`
	ConnectTimeout       int    `default:"5000"`
	PreferSimpleProtocol bool   `default:"true"`
	MaxOpenConns         int    `default:"100"`
	MaxIdleConns         int    `default:"1"`
}

func NewDB(config PGConfig) (*sqlx.DB, error) {
	dbConfig, err := pgx.ParseConfig(config.ConnectionString)
	if err != nil {
		return nil, err
	}

	dbConfig.ConnectTimeout = time.Duration(config.ConnectTimeout) * time.Millisecond
	dbConfig.PreferSimpleProtocol = config.PreferSimpleProtocol

	db := stdlib.OpenDB(*dbConfig)
	pg := sqlx.NewDb(db, "pgx")

	pg.SetMaxOpenConns(config.MaxOpenConns)
	pg.SetMaxIdleConns(config.MaxIdleConns)

	return pg, nil
}
