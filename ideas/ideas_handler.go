package ideas

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type NewIdea struct {
	Title       string
	Description string
	UserID      int64 `json:"user_id"`
}

type User struct {
	ID   int64
	Name string
}

func SaveIdeaHandler(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var idea NewIdea
		if err := json.NewDecoder(r.Body).Decode(&idea); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"parsing data failed"}`))
			return
		}
		if err := service.Save(r.Context(), idea); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"saving idea failed"}`))
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func FetchIdeasHandler(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		r.ParseForm()
		userID, err := strconv.ParseInt(r.FormValue("user_id"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"bad request with invalid user id"}`))
			return
		}

		ideas, err := service.Fetch(r.Context(), userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"fetching ideas failed"}`))
			return
		}

		if err := json.NewEncoder(w).Encode(ideas); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"decoding ideas failed"}`))
		}
	}
}
