package session

import (
	"context"
	"time"

	"github.com/gkits/gosanta/internal/session/entity"
)

type Reader interface {
	Get(ctx context.Context, token string) (*entity.Session, error)
}

type Writer interface {
	NewSession(ctx context.Context) (*entity.Session, error)
	Delete(ctx context.Context, token string) error
}

type Repository interface {
	Reader
	Writer
}

type Service struct {
	repo      Repository
	sessionCh chan *entity.Session
}

func New(repo Repository) *Service {
	service := &Service{
		repo:      repo,
		sessionCh: make(chan *entity.Session),
	}
	return service
}

func (service *Service) GetSession(ctx context.Context, token string) (*entity.Session, error) {
	session, err := service.repo.Get(ctx, token)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (service *Service) CreateNewSession(ctx context.Context) (*entity.Session, error) {
	session, err := service.repo.NewSession(ctx)
	if err != nil {
		return nil, err
	}

	time.AfterFunc(time.Second, func() {
		if err := service.repo.Delete(ctx, session.Token); err != nil {
			panic(err)
		}
	})

	service.sessionCh <- session

	return session, nil
}
