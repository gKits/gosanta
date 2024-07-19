package sqlite

import (
	"context"
	"database/sql"

	"github.com/gkits/gosanta/internal/santa/entity"
	_ "github.com/mattn/go-sqlite3"
)

type SantaSQLite struct {
	db *sql.DB
}

func New(db *sql.DB) (*SantaSQLite, error) {
	return &SantaSQLite{
		db: db,
	}, nil
}

func (repo *SantaSQLite) Get(ctx context.Context, id int64) (*entity.Player, error) {
	stmt, err := repo.db.PrepareContext(ctx, `SELECT * FROM Santa WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var player entity.Player
	if err := stmt.QueryRowContext(ctx, id).Scan(&player); err != nil {
		return nil, err
	}

	return &player, nil
}

func (repo *SantaSQLite) GetAll(ctx context.Context) ([]entity.Player, error) {
	stmt, err := repo.db.PrepareContext(ctx, `SELECT * FROM Santa`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	players := []entity.Player{}
	for rows.Next() {
		var player entity.Player
		if err := rows.Scan(&player); err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
}

func (repo *SantaSQLite) Insert(ctx context.Context, player entity.Player) (int64, error) {
	stmt, err := repo.db.PrepareContext(ctx, `INSERT INTO Santa(name, email) values(?, ?)`)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, player.Name, player.Email)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, err
}

func (repo *SantaSQLite) Update(ctx context.Context, player entity.Player) error {
	stmt, err := repo.db.PrepareContext(ctx, `UPDATE Santa SET name = ?, email = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, player.Name, player.Email, player.ID)
	if err != nil {
		return err
	}

	if affected, err := res.RowsAffected(); err != nil {
		return err
	} else if affected < 1 {
		return nil
	}

	return nil
}

func (repo *SantaSQLite) Delete(ctx context.Context, id int64) error {
	stmt, err := repo.db.PrepareContext(ctx, `DELETE FROM Santa WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	if affected, err := res.RowsAffected(); err != nil {
		return err
	} else if affected < 1 {
		return nil
	}

	return nil
}
