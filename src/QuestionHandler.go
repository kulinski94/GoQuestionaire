package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"fmt"

	"github.com/gorilla/mux"
)

func getAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions := RepoGetAllQuestions()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(questions); err != nil {
		panic(err)
	}
}

func addQuestion(w http.ResponseWriter, r *http.Request) {
	var question Question
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &question); err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	fmt.Println(question)

	t := RepoInsertQuestion(question)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Status", "200")
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func deleteQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionID, er := strconv.Atoi(vars["questionId"])

	if er != nil {
		panic(er)
	}

	res := RepoDeleteQuestion(questionID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
