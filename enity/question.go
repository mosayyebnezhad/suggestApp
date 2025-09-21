package entity

type Question struct {
	ID              uint
	Text            string
	PossibleAnswers []PossibleAnswers
	CorrectAnswerID uint
	Difficulty      QuestionDifficulty
	CategoryID      uint
}

type PossibleAnswers struct {
	ID     uint
	Text   string
	Choice PossibleAnswerChoice
}

type PossibleAnswerChoice uint8

func (p PossibleAnswerChoice) IsValid() bool {
	return p >= PossibleAnswerA && p <= PossibleAnswerD
}

/*
	 i write this but Hossein Nazari write this:
	func (p PossibleAnswerChoice) IsV alid() bool {
		if p >= 1 && p <= 4 {
			return true
		}
		return false
	}
*/

const (
	PossibleAnswerA PossibleAnswerChoice = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

type QuestionDifficulty uint8

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)

func (q QuestionDifficulty) IsValid() bool {
	return q >= QuestionDifficultyEasy && q <= QuestionDifficultyHard
}
