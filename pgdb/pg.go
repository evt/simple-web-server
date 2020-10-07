package pgdb

import (
	"time"

	"github.com/evt/rest-api-example/config"
	"github.com/go-pg/pg/v10"
)

// Timeout is a Postgres timeout
const Timeout = 5

// PgDB is a shortcut structure to a Postgres DB
type PgDB struct {
	*pg.DB
}

// Dial creates new database connection to postgres
func Dial(cfg *config.Config) (*PgDB, error) {
	pgOpts := &pg.Options{}

	var err error

	if cfg.PgURL == "" {
		return nil, nil
	}
	pgOpts, err = pg.ParseURL(cfg.PgURL)
	if err != nil {
		return nil, err
	}

	pgDB := pg.Connect(pgOpts)

	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	if Timeout > 0 {
		pgDB.WithTimeout(time.Second * time.Duration(Timeout))
	}

	return &PgDB{pgDB}, nil
}