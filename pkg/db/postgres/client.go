package postgres

import (
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

type Client interface {
	Begin() (*pgx.Tx, error)
	Query(sql string, args ...interface{}) (*pgx.Rows, error)
	QueryRow(sql string, args ...interface{}) *pgx.Row
}

func NewClient(cfg pgx.ConnConfig) Client {

	config := pgx.ConnPoolConfig{
		ConnConfig: cfg,
	}

	connection, err := pgx.NewConnPool(config)

	if err != nil {
		log.Println("Failed to connect to database")
	}

	return connection
}

func NewError(err error) error {
	var pgError *pgx.PgError
	if errors.As(err, &pgError) {
		return fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgError.Error(), pgError.Detail, pgError.Where, pgError.Code, pgError.SQLState())
	} else {
		return fmt.Errorf("unknown error: %s", err.Error())
	}
}
