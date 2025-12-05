package client

import (
	"context"
	"fmt"
	"log"

	"backend/config"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/Thiht/transactor"
	transactorpgx "github.com/Thiht/transactor/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBConnection(ctx context.Context, srvCfg *config.Server) *pgxpool.Pool {
	url := srvCfg.DBURL
	if url == "" {
		url = fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			srvCfg.DBUsername,
			srvCfg.DBPassword,
			srvCfg.DBHost,
			srvCfg.DBPort,
			srvCfg.DBName,
		)
	}

	conn, err := pgxpool.New(ctx, url)
	if err != nil {
		log.Printf("Cannot connect to Db: %v", err)
		return nil
	}
	return conn
}

func NewDBQueries(c *pgxpool.Pool) *sqlc.Queries {
	q := sqlc.New(c)
	return q
}

func NewDBTransactor(c *pgxpool.Pool) transactor.Transactor {
	t, _ := transactorpgx.NewTransactorFromPool(c)
	return t
}
