package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"github.com/weimpact/webother/logger"
)

type Server struct {
	Name string `default:"webother"`
}

type Store struct {
	Location string `required:"true"`
}

type Application struct {
	server Server
	db     DB
	store  Store
	port   int
}

var app Application

func Load() {
	var loadErr []error
	app.port, _ = strconv.Atoi(os.Getenv("PORT"))

	if err := envconfig.Process("APP", &app.server); err != nil {
		loadErr = append(loadErr, err)
	}
	if err := envconfig.Process("DB", &app.db); err != nil {
		loadErr = append(loadErr, err)
	}
	if err := envconfig.Process("STORE", &app.store); err != nil {
		loadErr = append(loadErr, err)
	}
	if len(loadErr) != 0 {
		var err string
		for _, e := range loadErr {
			err = fmt.Sprintf("%s%s", err, e)
		}
		logger.Errorf("%+v %+v", app, loadErr, err)
		panic(err)
	}
}

func AppPort() string {
	return fmt.Sprintf(":%d", app.port)
}

func Database() DB {
	return app.db
}

func StoreLocation() string {
	return app.store.Location
}

func StaticPort() string {
	return fmt.Sprintf(":%d", app.port)
}
