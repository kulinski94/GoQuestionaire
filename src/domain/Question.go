package main

import (
	"time"
)

type Question struct {
	Text    string    `json:"text"`
	Answers Answers   `json:"completed"`
	Due     time.Time `json:"due"`
}

type Answer struct {
}

type Answers []Answer

type Todos []Todo
