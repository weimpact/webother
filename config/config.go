package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Server struct {
	Port int `default:"8080"`
}

type Application struct {
	server Server
	db     DB
}

var app Application

func Load() {
	var loadErr []error

	if err := envconfig.Process("APP", &app.server); err != nil {
		loadErr = append(loadErr, err)
	}
	if err := envconfig.Process("DB", &app.db); err != nil {
		loadErr = append(loadErr, err)
	}
	fmt.Printf("%+v %+v", app, loadErr)
}

func AppPort() string {
	return fmt.Sprintf(":%d", app.server.Port)
}

func Database() DB {
	return app.db
}
