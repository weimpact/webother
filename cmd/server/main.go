package main

import (
	"net/http"

	"github.com/weimpact/webother/config"
)

func main() {
	config.Load()
	err := http.ListenAndServe(config.AppPort(), server())
	if err != nil {
		panic(err)
	}
}
