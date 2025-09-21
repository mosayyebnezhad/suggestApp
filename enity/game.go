package entity

import "time"

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	PlayerIDs   []uint
	StartTime   time.Time
}

type Player struct {
	ID      uint
	UserID  uint
	GameID  uint
	Score   uint
	Answers []PlayerAnswers
}

type PlayerAnswers struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
	Choice     PossibleAnswerChoice
}
