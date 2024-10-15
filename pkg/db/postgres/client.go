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
	Exec(sql string, arguments ...interface{}) (commandTag pgx.CommandTag, err error)
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
	if errors.Is(err, pgError) {
		e := fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgError.Error(), pgError.Detail, pgError.Where, pgError.Code, pgError.SQLState())
		log.Println(e)
		return fmt.Errorf(e)
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	e := fmt.Sprintf("unknown error: %s", err.Error())
	log.Println(e)
	return fmt.Errorf("unknown error: %s", err.Error())
}

func HandleTxCommit(tx *pgx.Tx) {
	if err := tx.Commit(); err != nil {
		log.Printf("failed to commit transaction: %s", err.Error())
	}
}
func HandleTxRollback(tx *pgx.Tx) {
	if err := tx.Rollback(); err != nil {
		log.Printf("failed to rollback transaction: %s", err.Error())
	}
}
