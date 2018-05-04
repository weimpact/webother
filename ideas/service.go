package ideas

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/weimpact/webother/logger"
)

type Idea struct {
	Title       string `db:"title"`
	Description string `db:"description"`
	UserID      int64  `db:"user_id"`
}

type Service struct {
	DB *sqlx.DB
}

func (s *Service) Save(ctx context.Context, i NewIdea) error {
	q := s.DB.Rebind("insert into ideas (user_id, title, description) values (?, ?, ?)")
	_, err := s.DB.Exec(q, i.UserID, i.Title, i.Description)
	if err != nil {
		logger.Errorf("[SaveIdea] error saving idea: %v with err: %v", i, err)
		return err
	}
	return nil
}

func (s *Service) Fetch(ctx context.Context, userID int64) ([]Idea, error) {
	var ideas []Idea
	q := s.DB.Rebind("select user_id, title, description from ideas where user_id = ?")
	err := s.DB.SelectContext(ctx, &ideas, q, userID)
	if err != nil {
		logger.Errorf("[SaveIdea] error fetching ideas for user: %v with err: %v", userID, err)
		return ideas, err
	}
	return ideas, nil
}
