package repo

import (
	"errors"
	"lib/app/match/models"
	"lib/common/id"
)

var ErrMatchAlreadyExists error = errors.New("match already exists")
var ErrMatchNotFound error = errors.New("match do not exists")

type MatchRepo interface {
	AddMatch(match models.Match) error

	GetUserMatch(userId id.ID) (models.Match, error)
	GetMatch(id id.ID) (models.Match, error)
	GetPublic() []models.Match
	GetAll() []models.Match

	RemoveMatch(matchId id.ID) error

	UpdateMatch(match models.Match) error
}

// i have a thought
// stop using interfaces and just think about implementing global methods (i think this is functional programming)
func GetMatchRepo() MatchRepo { return GetRepo() }

// also you can work on naming
// consider using func AddMatchRepo() IAddMatchRepo {}
