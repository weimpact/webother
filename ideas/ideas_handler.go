package ideas

import (
	"encoding/json"
	"net/http"
)

type Idea struct {
	Title       string
	Description string
	UserName    string
}

func SaveIdeaHandler(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var idea Idea
		if err := json.NewDecoder(r.Body).Decode(&idea); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err := service.Save(r.Context(), idea); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
