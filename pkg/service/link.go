package service

import (
	shorter "github.com/p12s/fintech-link-shorter"
	"github.com/p12s/fintech-link-shorter/pkg/repository"
)

type LinkService struct {
	repo repository.Link
}

func NewLinkService(repo repository.Link) *LinkService {
	return &LinkService{repo: repo}
}

func (l *LinkService) Create(longLink string) (shorter.UserLink, error) {
	result, err := l.repo.Create(longLink)
	if err != nil {
		return shorter.UserLink{}, err
	}

	return result, nil
}
