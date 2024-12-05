package usecases

import (
	"lib/app/match/messages"
	"lib/app/match/models"
	matchrepo "lib/app/match/repo"
	questionmodels "lib/app/questions/models"
	"lib/app/users/dto"
)

type AnswerArgs struct {
	Session dto.Session
	Answer  questionmodels.AnswerMessage `json:"answer"`
}

func Answer(args AnswerArgs) error {
	matchRepo := matchrepo.GetMatchRepo()
	match, err := matchRepo.GetUserMatch(args.Session.UserId)
	if err != nil {
		return err
	}

	if err := match.Answer(args.Session.UserId, args.Answer); err != nil {
		return err
	}

	question, _ := match.GetCurrentQuestion()
	if question.AnsweredCorrectly != nil && *question.AnsweredCorrectly {
		match.GetUser(args.Session.UserId).Score += 1000
	}

	messages.SendScores(match)
	if err := match.NextQuestion(); err == models.ErrNoQuestionsLeft {
		messages.SendMatch(match)
	} else {
		messages.SendQuestion(match)
		EnsureAnswered(EnsureAnswredArgs{Match: match})
	}
	if err := matchRepo.UpdateMatch(match); err != nil {
		return err
	}

	return nil
}
