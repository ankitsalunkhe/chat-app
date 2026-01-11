package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ankitsalunkhe/chat-app-backend/internal/storage/migrations"
	"github.com/lib/pq"
)

var (
	ErrNoSuchUsername    = errors.New("no such username exists")
	ErrDuplicateUsername = errors.New("username already exists")
)

type Database struct {
	db *sql.DB
}

type User struct {
	Id       int32  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func New(connStr string) (Database, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return Database{}, fmt.Errorf("unable to open database: %w", err)
	}

	if err := migrations.New(db); err != nil {
		return Database{}, fmt.Errorf("running migrations: %w", err)
	}

	return Database{db: db}, nil
}

func (db *Database) GetUser(username string) (User, error) {
	rows := db.db.QueryRow(`SELECT * FROM chat_app.credentials WHERE username = $1`, username)

	var user User
	if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, ErrNoSuchUsername
		}
		return user, fmt.Errorf("scanning database result : %w", err)
	}

	return user, nil
}

func (db *Database) CreateCredentials(username string, password string) error {
	tx, err := db.db.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO chat_app.credentials (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		err, ok := err.(*pq.Error)
		if !ok {
			return fmt.Errorf("executing transaction: %w", err)
		}

		if err.Code.Name() == "unique_violation" {
			return ErrDuplicateUsername
		}

		return fmt.Errorf("executing transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	id, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("fetching rows affected: %w", err)
	}

	if id == 0 {
		return fmt.Errorf("no rows added")
	}

	return nil
}
