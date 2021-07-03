package service

import (
	shorter "github.com/p12s/fintech-link-shorter"
	"github.com/p12s/fintech-link-shorter/pkg/repository"
)

// LinkService - второй уровень "луковой архитектуры"
// он передает действие репозиторию
type LinkService struct {
	repo repository.Link
}

// NewLinkService - конструктор
func NewLinkService(repo repository.Link) *LinkService {
	return &LinkService{repo: repo}
}

// Create - создание короткой ссылки
func (l *LinkService) Create(longLink string) (shorter.UserLink, error) {
	result, err := l.repo.Create(longLink)
	if err != nil {
		return shorter.UserLink{}, err
	}

	return result, nil
}

// Long - получение длинной ссылки
func (l *LinkService) Long(shortLink string) (shorter.UserLink, error) {
	result, err := l.repo.Long(shortLink)
	if err != nil {
		return shorter.UserLink{}, err
	}

	return result, nil
}
