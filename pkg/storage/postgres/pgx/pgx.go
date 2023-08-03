package pgx

import (
	"context"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"marketplace/internal/config"
	"marketplace/pkg/storage/postgres"
)

type client struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewClient(ctx context.Context, dbConf config.Database, maxConns int32) (postgres.Database, error) {
	poolConfig, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Println("pgxpool.ParseConfig ERROR: " + err.Error())
		return nil, errors.New("pgxpool.ParseConfig ERROR: " + err.Error())
	}
	poolConfig.ConnConfig.Host = dbConf.Host
	poolConfig.ConnConfig.Port = dbConf.Port
	poolConfig.ConnConfig.User = dbConf.User
	poolConfig.ConnConfig.Password = dbConf.Password
	poolConfig.ConnConfig.Database = dbConf.Name
	//poolConfig.AfterConnect = afterConnect
	poolConfig.MaxConns = maxConns
	poolConfig.MaxConns = 10

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		log.Println("pgxpool.ConnectConfig ERROR: " + err.Error())
		return nil, errors.New("pgxpool.ConnectConfig ERROR: " + err.Error())
	}
	return &client{pool: pool, ctx: ctx}, nil
}

func (c *client) Close() {
	c.pool.Close()
}

func (c *client) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return c.pool.Query(ctx, sql, args...)
}

func (c *client) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return c.pool.QueryRow(ctx, sql, args...)
}

func (c *client) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return c.pool.Exec(ctx, sql, arguments...)
}
