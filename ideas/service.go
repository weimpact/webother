package ideas

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/weimpact/webother/logger"
)

type Service struct {
	DB *sqlx.DB
}

func (s *Service) Save(ctx context.Context, i Idea) error {
	q := s.DB.Rebind("insert into ideas (user_id, title, description) values (?, ?, ?)")
	_, err := s.DB.Exec(q, i.UserID, i.Title, i.Description)
	if err != nil {
		logger.Errorf("[SaveIdea] error saving idea: %v with err: %v", i, err)
		return err
	}
	return nil
}
