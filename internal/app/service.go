package app

import (
	log "github.com/sirupsen/logrus"

	"github.com/codebind-luna/booking-service/internal/domain"
)

type ticketService struct {
	logger *log.Logger
	repo   domain.Repository
}

func NewService(logger *log.Logger, repo domain.Repository) *ticketService {
	return &ticketService{
		logger: logger,
		repo:   repo,
	}
}
