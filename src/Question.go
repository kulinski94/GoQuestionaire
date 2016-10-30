package main

type Question struct {
	Id      int     `json:"id"`
	Text    string  `json:"text"`
	Answers Answers `json:"answers"`
}

type Answer struct {
	Id      int    `json:"id"`
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

type Answers []Answer

type Questions []Question
