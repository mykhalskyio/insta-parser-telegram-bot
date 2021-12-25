package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mykhalskyio/insta-parser-telegram-bot/internal/config"
	_ "github.com/mykhalskyio/insta-parser-telegram-bot/internal/migrations"
	"github.com/pressly/goose/v3"
)

// db postgres struct
type Postgres struct {
	db *sql.DB
}

// create postgres struct
func NewConnect(cfg *config.Config) (*Postgres, error) {
	conect, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=%s ", cfg.Postgres.Name, cfg.Postgres.User, cfg.Postgres.Pass, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Sslmode))
	if err != nil {
		return nil, err
	}

	log.Println("Connect bd postgres - OK")

	if err = conect.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{db: conect}, nil
}

// check result
func (pg *Postgres) Check(id string) bool {
	var result bool
	pg.db.QueryRow("SELECT result FROM parser WHERE id_storis = $1", id).Scan(&result)

	if result {
		return true
	} else {
		pg.db.Exec("INSERT INTO parser(id_storis, result) VALUES($1, true)", id)
	}

	return false
}

func (pg *Postgres) CheckTable() error {
	err := goose.Up(pg.db, ".")
	if err != nil {
		return err
	}

	return nil
}
