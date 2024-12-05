package models

import (
	"errors"
	"lib/app/questions/models"
	"lib/common/id"
	"time"

	"golang.org/x/exp/rand"
)

var ErrCannotModifyInMatch error = errors.New("cannot modify match settings when playing")
var ErrMatchHasToStart error = errors.New("match hasn't started")
var ErrMatchHasStarted error = errors.New("match has already started")
var ErrLimitCannotBeBelowMatchLen error = errors.New("limit cannot be below match length")
var ErrServerIsFull = errors.New("server is full")
var ErrAlreadyInMatch error = errors.New("already in match")
var ErrNotAHost error = errors.New("you are not a host")
var ErrQuestionExists error = errors.New("question already exists")
var ErrQuestionNotFound error = errors.New("do not found a question")
var ErrNoQuestionsLeft error = errors.New("all questions got answered")
var ErrToFewQuestions error = errors.New("to few questions to start match")

type Match struct {
	MatchId         id.ID
	Players         []PlayerData
	PlayersLimit    PlayersLimit
	IsPublic        bool
	Questions       []QuestionData
	CurrentQuestion int
	StartedAt       *time.Time
}

func NewMatch(host PlayerData) (Match, error) {
	if !host.IsHost {
		return Match{}, errors.New("player is not an admin")
	}

	return Match{ // are this comments required ? isn't that obvious ?
		MatchId:         id.New(),
		Players:         []PlayerData{host}, // only host is in new match
		PlayersLimit:    2,                  // default player limit
		IsPublic:        false,              // by default isn't public
		Questions:       []QuestionData{},   // no questions by default
		CurrentQuestion: -1,                 //
		StartedAt:       nil,                // match hasn't started
	}, nil
}

func (match *Match) CanStart() error {
	if match.Started() {
		return ErrMatchHasStarted
	}
	if len(match.Questions) < 1 {
		return ErrToFewQuestions
	}
	return nil
}

func (match *Match) Start() error {
	if err := match.CanStart(); err != nil {
		return err
	}
	now := time.Now()
	match.StartedAt = &now
	rand.Seed(uint64(now.UnixNano()))
	rand.Shuffle(len(match.Questions), func(i, j int) {
		match.Questions[i], match.Questions[j] = match.Questions[j], match.Questions[i]
	})
	match.NextQuestion()
	return nil
}

func (match *Match) Started() bool {
	return match.StartedAt != nil
}

func (match *Match) SetPublic(public bool, userId id.ID) error {
	if !match.IsHost(userId) {
		return ErrNotAHost
	}
	if match.Started() {
		return ErrCannotModifyInMatch
	}
	match.IsPublic = public
	return nil
}

func (match *Match) SetPlayersLimit(limit PlayersLimit, userId id.ID) error {
	if !match.IsHost(userId) {
		return ErrNotAHost
	}
	if match.Started() {
		return ErrCannotModifyInMatch
	}
	if len(match.Players) > int(limit) {
		return ErrLimitCannotBeBelowMatchLen
	}
	match.PlayersLimit = limit
	return nil
}

func (match *Match) Reset() error {
	if !match.Started() {
		return ErrMatchHasToStart
	}
	for i := range match.Players {
		match.Players[i].Reset()
	}
	for i := range match.Questions {
		match.Questions[i].Reset()
	}
	match.StartedAt = nil
	match.CurrentQuestion = -1
	return nil
}

func (match *Match) Join(player PlayerData) error {
	if len(match.Players) == int(match.PlayersLimit) {
		return ErrServerIsFull
	}
	for _, playerInMatch := range match.Players {
		if playerInMatch.UserId == player.UserId {
			return ErrAlreadyInMatch
		}
	}
	match.Players = append(match.Players, player)
	return nil
}

func (match *Match) Quit(userId id.ID) error {
	isHost := false
	userIndex := -1
	for i, player := range match.Players {
		if player.UserId == userId {
			userIndex = i
			isHost = player.IsHost
			break
		}
	}

	if userIndex == -1 {
		return errors.New("do not found user")
	}

	match.Players = append(match.Players[:userIndex], match.Players[userIndex+1:]...)

	if isHost && len(match.Players) != 0 {
		match.Players[0].IsHost = isHost
	}

	return nil
}

func (match *Match) IsHost(userId id.ID) bool {
	for _, player := range match.Players {
		if player.UserId == userId {
			return player.IsHost
		}
	}

	return false
}

func (match *Match) AddQuestion(question models.Question, userId id.ID) error {
	if !match.IsHost(userId) {
		return ErrNotAHost
	}
	for _, existingQuestion := range match.Questions {
		if existingQuestion.QuestionId == question.QuestionId {
			return ErrQuestionExists
		}
	}
	match.Questions = append(match.Questions, NewQuestion(question))
	return nil
}

func (match *Match) RemoveQuestion(questionId id.ID, userId id.ID) error {
	if !match.IsHost(userId) {
		return ErrNotAHost
	}

	questionIndex := -1
	for i, question := range match.Questions {
		if question.QuestionId == questionId {
			questionIndex = i
			break
		}
	}
	if questionIndex == -1 {
		return ErrQuestionNotFound
	}

	match.Questions = append(match.Questions[:questionIndex], match.Questions[questionIndex+1:]...)
	return nil
}

func (match *Match) NextQuestion() error {
	if !match.Started() {
		return ErrMatchHasToStart
	}
	if match.CurrentQuestion+1 >= len(match.Questions) {
		match.Reset()
		return ErrNoQuestionsLeft
	}
	time.Sleep(3 * time.Second)
	match.CurrentQuestion += 1
	return nil
}

func (match *Match) GetUser(userId id.ID) *PlayerData {
	for i, player := range match.Players {
		if player.User.UserId == userId {
			return &match.Players[i]
		}
	}
	return nil
}

func (match *Match) GetCurrentQuestion() (*QuestionData, error) {
	if !match.Started() {
		return nil, ErrMatchHasToStart
	}
	return &match.Questions[match.CurrentQuestion], nil
}

func (match *Match) Answer(userId id.ID, answer models.AnswerMessage) error {
	if !match.Started() {
		return ErrMatchHasStarted
	}
	question, _ := match.GetCurrentQuestion()
	if err := question.Answer(userId, answer); err != nil {
		return err
	}

	return nil
}
