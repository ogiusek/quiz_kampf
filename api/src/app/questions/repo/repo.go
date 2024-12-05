package repo

import (
	"errors"
	"lib/app/questions/models"
	"lib/common/id"
	"lib/services"
)

type QuestionRepo interface {
	GetQuestion(id.ID) *models.Question
	AddQuestion(models.Question) error
	RemoveQuestion(questionId id.ID) error
	UpdateQuestion(question models.Question) error
	Search(SearchQuestionsArgs) []models.Question
	UserQuestions(userId id.ID, page Page) []models.Question
}

var ErrQuestionAlreadyExist error = errors.New("this id is already taken")
var ErrQuestionDoesNotExist error = errors.New("question does not exist")

type questionRepo struct{}

func (questionRepo) GetQuestion(questionId id.ID) *models.Question {
	db, free := services.Db()
	defer free()
	var question models.Question
	db.
		Where("question_id = ?", questionId).
		Preload("Creator").
		Find(&question)
	return &question
}

func (questionRepo) AddQuestion(question models.Question) error {
	db, free := services.Db()
	defer free()
	if err := db.Create(question).Error; err != nil {
		return ErrQuestionAlreadyExist
	}
	return nil
}

func (questionRepo) RemoveQuestion(questionId id.ID) error {
	db, free := services.Db()
	defer free()
	tx := db.Where("question_id = ?", questionId).
		Unscoped().
		Delete(&models.Question{})

	if tx.RowsAffected == 0 {
		return ErrQuestionDoesNotExist
	}
	return nil
}

func (questionRepo) UpdateQuestion(question models.Question) error {
	db, free := services.Db()
	defer free()
	db.Save(&question)
	return nil
}

func (questionRepo) UserQuestions(userId id.ID, page Page) []models.Question {
	db, free := services.Db()
	defer free()

	var questions []models.Question
	db.
		Where("creator_id = ?", userId).
		Preload("Creator").
		Limit(20).
		Offset(20 * int(page)).
		Find(&questions)
	return questions
}

func GetQuestionRepo() QuestionRepo {
	return questionRepo{}
}
