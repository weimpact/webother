package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/weimpact/webother/logger"
)

type Server struct {
	Port int `default:"8080"`
}

type Store struct {
	Location string `required:"true"`
	Port     int    `default:"9090"`
}

type Application struct {
	server Server
	db     DB
	store  Store
}

var app Application

func Load() {
	var loadErr []error

	if err := envconfig.Process("", &app.server); err != nil {
		loadErr = append(loadErr, err)
	}
	if err := envconfig.Process("DB", &app.db); err != nil {
		loadErr = append(loadErr, err)
	}
	if err := envconfig.Process("STORE", &app.store); err != nil {
		loadErr = append(loadErr, err)
	}
	logger.Errorf("%+v %+v", app, loadErr)
}

func AppPort() string {
	return fmt.Sprintf(":%d", app.server.Port)
}

func Database() DB {
	return app.db
}

func StoreLocation() string {
	return app.store.Location
}

func StaticPort() string {
	return fmt.Sprintf(":%d", app.store.Port)
}
