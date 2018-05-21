package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/weimpact/webother/config"
	"github.com/weimpact/webother/logger"
)

func main() {
	config.Load()
	var wg sync.WaitGroup
	wg.Add(2)
	go start(server(), config.AppPort(), &wg)
	go start(static(), config.StaticPort(), &wg)
	wg.Wait()
}

func start(s *mux.Router, addr string, wg *sync.WaitGroup) {
	defer wg.Done()
	var server http.Server
	go func() {
		server = http.Server{Addr: addr, Handler: s}
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("started server in addr: %s", addr)
	<-killer()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	logger.Infof("stopped server in addr %s", addr)
}

func killer() <-chan os.Signal {
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, os.Interrupt, os.Kill)
	return kill
}
