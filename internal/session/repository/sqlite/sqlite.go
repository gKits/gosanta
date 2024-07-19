package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/gkits/gosanta/internal/session/entity"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SessionSQLite struct {
	db *sql.DB
}

func New(db *sql.DB) (*SessionSQLite, error) {
	return &SessionSQLite{
		db: db,
	}, nil
}

func (repo *SessionSQLite) Get(ctx context.Context, token string) (*entity.Session, error) {
	stmt, err := repo.db.PrepareContext(ctx, `SELECT * FROM Sessions WHERE token = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var player entity.Session
	if err := stmt.QueryRowContext(ctx, token).Scan(&player); err != nil {
		return nil, err
	}

	return &player, nil
}

func (repo *SessionSQLite) NewSession(ctx context.Context) (*entity.Session, error) {
	session := entity.Session{
		Token:     uuid.NewString(),
		CreatedAt: time.Now(),
	}

	stmt, err := repo.db.PrepareContext(ctx, `INSERT INTO Sessions(token, created_at) values(?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, session.Token, session.CreatedAt)
	if err != nil {
		return nil, err
	}

	if affected, err := res.RowsAffected(); err != nil {
		return nil, err
	} else if affected < 1 {
		return nil, sql.ErrNoRows
	}

	return &session, nil
}

func (repo *SessionSQLite) Delete(ctx context.Context, token string) error {
	stmt, err := repo.db.PrepareContext(ctx, `DELETE FROM Sessions WHERE token = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, token)
	if err != nil {
		return err
	}

	if affected, err := res.RowsAffected(); err != nil {
		return err
	} else if affected < 1 {
		return sql.ErrNoRows
	}

	return nil
}
