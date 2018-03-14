package ideas

import (
	"context"
	"database/sql"
)

type Service struct {
	db sql.DB
}

func (s *Service) Save(ctx context.Context, i Idea) error {
}
