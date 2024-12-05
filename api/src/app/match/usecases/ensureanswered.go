package usecases

import (
	"fmt"
	"lib/app/match/messages"
	"lib/app/match/models"
	"lib/app/match/repo"
	"time"
)

type EnsureAnswredArgs struct {
	Match models.Match
}

func EnsureAnswered(args EnsureAnswredArgs) { // this has no response type because it is called by other use cases and not by websocket
	go func() {
		matchRepo := repo.GetMatchRepo()
		currentQuestionIndex := args.Match.CurrentQuestion
		currentQuestion, err := args.Match.GetCurrentQuestion()
		if err != nil {
			panic(fmt.Sprintf("error when getting question in 'EnsureAnswered' method %v", err))
		}

		time.Sleep(time.Duration(currentQuestion.AnswerTimeInSeconds) * time.Second)
		args.Match, err = matchRepo.GetMatch(args.Match.MatchId)
		if err != nil {
			return
		}
		if currentQuestionIndex != args.Match.CurrentQuestion {
			return
		}

		messages.SendScores(args.Match)
		if err := args.Match.NextQuestion(); err == models.ErrNoQuestionsLeft {
			messages.SendMatch(args.Match)
			matchRepo.UpdateMatch(args.Match)
			return
		}
		messages.SendQuestion(args.Match)
		if err := matchRepo.UpdateMatch(args.Match); err != nil {
			return
		}
		EnsureAnswered(EnsureAnswredArgs{Match: args.Match})
	}()
}
