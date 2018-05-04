package files

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/weimpact/webother/config"
)

type Service struct {
	DB *sqlx.DB
}

func (fs *Service) Save(ctx context.Context, d Data) error {
	location, err := fs.SaveToDisk(ctx, d)
	if err != nil {
		return err
	}
	q := fs.DB.Rebind("insert into files (user_id, location) values (?, ?)")
	_, err = fs.DB.ExecContext(ctx, q, d.userID, location)
	if err != nil {
		return err
	}
	return nil
}

func (fs *Service) SaveToDisk(ctx context.Context, d Data) (string, error) {
	id := uuid.NewV4()
	location := fmt.Sprintf("%s/%s", config.StoreLocation(), id.String())
	f, err := os.Create(location)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, d.file)
	if err != nil {
		return "", err
	}
	return location, nil
}
