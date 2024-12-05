package repo

import (
	"lib/app/match/models"
	"lib/common/id"
	"sync"
)

type Repo struct {
	Matches    map[id.ID]models.Match
	writeMutex sync.Mutex
}

func (repo *Repo) GetUserMatch(userId id.ID) (models.Match, error) {
	for _, match := range repo.Matches {
		for _, player := range match.Players {
			if player.UserId == userId {
				return match, nil
			}
		}
	}
	return models.Match{}, ErrMatchNotFound
}

func (repo *Repo) GetMatch(matchId id.ID) (models.Match, error) {
	match, ok := repo.Matches[matchId]
	if !ok {
		return match, ErrMatchNotFound
	}
	return match, nil
}

func (repo *Repo) GetPublic() []models.Match {
	res := []models.Match{}
	for _, match := range repo.Matches {
		if match.IsPublic {
			res = append(res, match)
		}
	}
	return res
}

func (repo *Repo) GetAll() []models.Match {
	res := []models.Match{}
	for _, match := range repo.Matches {
		res = append(res, match)
	}
	return res
}
func (repo *Repo) AddMatch(match models.Match) error {
	repo.writeMutex.Lock()
	_, ok := repo.Matches[match.MatchId]
	if ok {
		return ErrMatchAlreadyExists
	}
	repo.Matches[match.MatchId] = match
	repo.writeMutex.Unlock()
	return nil
}

func (repo *Repo) UpdateMatch(match models.Match) error {
	repo.writeMutex.Lock()
	_, ok := repo.Matches[match.MatchId]
	if !ok {
		return ErrMatchNotFound
	}
	repo.Matches[match.MatchId] = match
	repo.writeMutex.Unlock()
	return nil
}

func (repo *Repo) RemoveMatch(matchId id.ID) error {
	repo.writeMutex.Lock()
	_, ok := repo.Matches[matchId]
	if !ok {
		return ErrMatchNotFound
	}
	delete(repo.Matches, matchId)
	repo.writeMutex.Unlock()
	return nil
}

var repo *Repo = &Repo{Matches: map[id.ID]models.Match{}}

func GetRepo() *Repo {
	return repo
}
