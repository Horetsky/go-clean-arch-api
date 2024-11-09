package postgres

import (
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
