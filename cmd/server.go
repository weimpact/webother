package cmd

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/weimpact/webother/ideas"
)

func main() {
	m := mux.NewRouter()
	m.HandlerFunc("/ideas", ideas.SaveIdeaHandler()).Methods(http.MethodPut)
}
