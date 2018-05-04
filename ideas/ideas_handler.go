package ideas

import (
	"encoding/json"
	"net/http"
)

type Idea struct {
	Title       string
	Description string
	UserID      int64 `json:"user_id"`
}

func SaveIdeaHandler(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var idea Idea
		if err := json.NewDecoder(r.Body).Decode(&idea); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"parsing data failed"}`))
			return
		}
		if err := service.Save(r.Context(), idea); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"saving idea failed"}`))
		}
	}
}
