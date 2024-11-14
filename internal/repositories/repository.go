package repositories

import (
	"errors"

	"github.com/codebind-luna/booking-service/internal/domain"
	inmemoryrepository "github.com/codebind-luna/booking-service/internal/repositories/in.memory"
	log "github.com/sirupsen/logrus"
)

// New - retrieve a repository
func New(
	logger *log.Logger,
	repoType domain.RepositoryType,
) (domain.Repository, error) {
	switch repoType {
	case domain.InMemoryRepository:
		return inmemoryrepository.NewInMemoryRepository(10), nil
	default:
		return nil, errors.New("invalid repository provided")
	}
}
