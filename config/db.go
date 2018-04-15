package config

import (
	"fmt"
	"time"
)

type DB struct {
	Driver            string `default:"postgres"`
	Host              string `required:"true"`
	User              string `required:"true"`
	Password          string `required:"true"`
	Port              int    `required:"true"`
	MaxIdleConns      int    `split_words:"true" default:"20"`
	MaxOpenConns      int    `split_words:"true" default:"30"`
	MaxConnLifetimeMs int    `split_words:"true" default:"1000"`
	Name              string `split_words:"true" required:"true"`
	SslMode           string `split_words:"true" default:"disable"`
}

func (db DB) MaxConnLifetime() time.Duration {
	return time.Millisecond * time.Duration(db.MaxConnLifetimeMs)
}

func (db DB) URL() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", db.User, db.Password, db.Host, db.Port, db.Name, db.SslMode)
}
