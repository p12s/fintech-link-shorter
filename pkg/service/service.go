package service

import (
	"github.com/p12s/fintech-link-shorter"
	"github.com/p12s/fintech-link-shorter/pkg/repository"
)

// Link - интерфейс для сервисов
type Link interface {
	Create(longLink string) (shorter.UserLink, error)
	Long(shortLink string) (shorter.UserLink, error)
}

// Service - сервис
type Service struct {
	Link
}

// NewService - конструктор
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Link: NewLinkService(repo.Link),
	}
}
