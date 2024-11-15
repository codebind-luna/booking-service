package domain

import (
	"fmt"

	"github.com/codebind-luna/booking-service/internal/domain/models"
)

type RepositoryType string

func (r RepositoryType) String() string {
	return string(r)
}

const (
	InMemoryRepository RepositoryType = "in-memory"
)

var (
	validRepositories = []RepositoryType{
		InMemoryRepository,
	}
	ErrInvalidRepositoryType = fmt.Errorf("invalid repository type")
	repositoryMap            = map[string]RepositoryType{
		InMemoryRepository.String(): InMemoryRepository,
	}
)

func isValidRepository(repo RepositoryType) bool {
	for _, valid := range validRepositories {
		if valid == repo {
			return true
		}
	}
	return false
}

func ParseRepository(repo string) (RepositoryType, error) {
	r, ok := repositoryMap[repo]
	if !ok {
		return "", ErrInvalidRepositoryType
	}
	if !isValidRepository(r) {
		return "", ErrInvalidRepositoryType
	}
	return r, nil
}

type Repository interface {
	CreateBooking(user *models.User, fromCity string, toCity string, price float64) (*string, error)
	GetReceiptByUser(email string) (*models.Ticket, error)
	GetSeatMapBySection(section string) ([]*models.Seat, error)
	ModifySeat(email string, section string, seatNo int) error
	RemoveUser(email string) error
}
