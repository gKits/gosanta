package santa

import (
	"context"

	"github.com/gkits/gosanta/internal/santa/entity"
)

type Reader interface {
	Get(ctx context.Context, id int64) (*entity.Player, error)
	GetAll(ctx context.Context) ([]entity.Player, error)
}

type Writer interface {
	Insert(ctx context.Context, player entity.Player) (int64, error)
	Update(ctx context.Context, player entity.Player) error
	Delete(ctx context.Context, id int64) error
}

type Repository interface {
	Reader
	Writer
}

type Service struct {
}

func GetPlayer(ctx context.Context, id int64) (*entity.Player, error) {
	return nil, nil
}

func GetPlayers(ctx context.Context) ([]entity.Player, error) {
	return nil, nil
}

func AddPlayer(ctx context.Context, player entity.Player) (int64, error) {
	return 0, nil
}

func EditPlayer(ctx context.Context, player entity.Player) error {
	return nil
}

func DeletePlayer(ctx context.Context, id int64) error {
	return nil
}
