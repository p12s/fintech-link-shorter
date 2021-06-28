package service

import (
	shorter "github.com/p12s/fintech-link-shorter"
	"github.com/p12s/fintech-link-shorter/pkg/repository"
)

type Link interface {
	Create(longLink string) (shorter.UserLink, error)
}

type Service struct {
	Link
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Link: NewLinkService(repo.Link),
	}
}
