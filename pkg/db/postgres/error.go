package postgres

import (
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

func NewError(err error) error {
	var pgError *pgx.PgError
	if errors.Is(err, pgError) {
		e := fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
			pgError.Error(), pgError.Detail, pgError.Where, pgError.Code, pgError.SQLState(),
		)
		log.Println(e)
		return fmt.Errorf(e)
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	e := fmt.Sprintf("unknown error: %s", err.Error())

	log.Println(e)
	return fmt.Errorf(e)
}
