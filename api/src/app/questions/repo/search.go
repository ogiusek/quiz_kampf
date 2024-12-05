package repo

import (
	"errors"
	"fmt"
	"lib/app/questions/models"
	"lib/services"
	"strings"
)

type Limit int

func (vo Limit) Valid() error {
	if int(vo) < 5 || int(vo) > 50 {
		return errors.New("limit has to be between 5 and 50")
	}
	return nil
}

type Page int

func (vo Page) Valid() error {
	if int(vo) < 0 {
		return errors.New("page cannot be below zero")
	}
	return nil
}

type SearchQuestionsArgs struct {
	SearchQuestion              string
	SearchQuestionCaseSensitive bool
	SearchAnswer                string
	SearchAnswerCaseSensitive   bool
	SearchNick                  string
	SearchNickCaseSentitive     bool
	Limit                       Limit
	Page                        Page
}

func (questionRepo) Search(args SearchQuestionsArgs) []models.Question {
	db, free := services.Db()
	defer free()

	var questions []models.Question

	var conditions []string = []string{}
	var conditionsArgs []any = []any{}

	if args.SearchQuestion != "" {
		if args.SearchQuestionCaseSensitive {
			conditions = append(conditions, "(question_text LIKE ?)")
		} else {
			conditions = append(conditions, "(UPPER(question_text) LIKE UPPER(?))")
		}
		conditionsArgs = append(conditionsArgs, "%"+args.SearchQuestion+"%")
	}

	if args.SearchAnswer != "" {
		var prefix string = ""
		if !args.SearchAnswerCaseSensitive {
			prefix = "UPPER"
		}
		conditions = append(conditions, fmt.Sprintf(`EXISTS (
			SELECT 1 
			FROM jsonb_each_text(answer) AS kv(key, value)
			WHERE %s(value) LIKE %s(?)
			)`, prefix, prefix))
		conditionsArgs = append(conditionsArgs, "%"+args.SearchAnswer+"%")
	}

	if args.SearchNick != "" {
		var prefix string = ""
		if !args.SearchNickCaseSentitive {
			prefix = "UPPER"
		}
		conditions = append(conditions, fmt.Sprintf(`(%s(users.user_name) LIKE %s(?))`, prefix, prefix))
		conditionsArgs = append(conditionsArgs, "%"+args.SearchNick+"%")
	}

	db.
		Preload("Creator").
		Where(strings.Join(conditions, " AND "), conditionsArgs...).
		Limit(int(args.Limit)).
		Offset(int(args.Limit) * int(args.Page)).
		Find(&questions)

	return questions
}
