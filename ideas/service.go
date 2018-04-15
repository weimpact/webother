package ideas

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB *sqlx.DB
}

func (s *Service) Save(ctx context.Context, i Idea) error {
	return nil
}
