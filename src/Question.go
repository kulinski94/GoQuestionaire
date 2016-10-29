package main

type Question struct {
	Text    string  `json:"text"`
	Answers Answers `json:"completed"`
}

type Answer struct {
	Text    string `json:"text"`
	Correct bool   `json:"correct"`
}

type Answers []Answer

type Questions []Question
