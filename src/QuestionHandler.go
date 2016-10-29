package main

import (
	"encoding/json"
	"net/http"
)

func getAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions := Questions{
		Question{Text: "Write presentation"},
		Question{Text: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(questions); err != nil {
		panic(err)
	}
}
