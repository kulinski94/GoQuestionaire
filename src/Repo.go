package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func RepoFindQuestion(id int) Question {

	return Question{}
}

func RepoGetAllQuestions() Questions {
	db, err := sql.Open("mysql", "go:go@/go_questionaire?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT id,text FROM questions")
	checkErr(err)

	var questions Questions

	for rows.Next() {
		var id int
		var text string
		err = rows.Scan(&id, &text)
		checkErr(err)
		q := Question{Id: id, Text: text, Answers: findAnswersForQuestion(id)}
		questions = append(questions, q)
	}
	return questions
}

func findAnswersForQuestion(QuestionId int) Answers {
	db, err := sql.Open("mysql", "go:go@/go_questionaire?charset=utf8")
	checkErr(err)

	answerRows, err := db.Query("SELECT id,text FROM answers where question_id = ?", QuestionId)
	checkErr(err)

	var answers Answers

	for answerRows.Next() {
		var id int
		var text string
		err = answerRows.Scan(&id, &text)
		checkErr(err)
		a := Answer{Id: id, Text: text}
		answers = append(answers, a)
	}
	return answers
}

func findQuestionId(Text string) int {

	db, err := sql.Open("mysql", "go:go@/go_questionaire?charset=utf8")
	checkErr(err)

	answerRow := db.QueryRow("SELECT id FROM questions where text = ?", Text)
	checkErr(err)

	var id int
	err = answerRow.Scan(&id)

	checkErr(err)

	return id
}

func RepoInsertQuestion(QuestionToAdd Question) sql.Result {
	db, err := sql.Open("mysql", "go:go@/go_questionaire?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT questions SET text=?")
	checkErr(err)

	res, err := stmt.Exec(QuestionToAdd.Text)
	checkErr(err)

	questionID := findQuestionId(QuestionToAdd.Text)

	for _, answer := range QuestionToAdd.Answers {
		RepoInsertAnswer(answer, questionID)
	}

	return res
}

func RepoInsertAnswer(Answer Answer, QuestionId int) sql.Result {
	db, err := sql.Open("mysql", "go:go@/go_questionaire?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT answers SET text=?, correct=?, question_id=?")
	checkErr(err)

	res, err := stmt.Exec(Answer.Text, Answer.Correct, QuestionId)
	checkErr(err)
	return res
}

func RepoDeleteQuestion(questionID int) sql.Result {
	db, err := sql.Open("mysql", "go:go@/go_questionaire?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("delete from questions where id=?")
	checkErr(err)

	res, err := stmt.Exec(questionID)
	checkErr(err)
	return res
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
