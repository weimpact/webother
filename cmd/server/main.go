package main

import (
	"net/http"

	"github.com/weimpact/webother/config"
)

func main() {
	config.Load()
	http.ListenAndServe(config.AppPort(), server())
}
