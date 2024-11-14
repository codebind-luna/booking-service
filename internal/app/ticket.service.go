package app

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/codebind-luna/booking-service/internal/domain"
	"github.com/codebind-luna/booking-service/internal/domain/models"
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

func (s *ticketService) PurchaseTicket(ctx context.Context, email, firstName, lastName, fromCity, toCity string, price float64) (*string, error) {
	user := models.NewUser(email, firstName, lastName)

	bookingID, bookingErr := s.repo.CreateBooking(user, fromCity, toCity, price)

	if bookingErr != nil {
		return nil, bookingErr
	}

	return bookingID, nil
}

func (s *ticketService) GetReceipt(ctx context.Context, email string) (*models.Ticket, error) {
	ticket, err := s.repo.GetReceiptByUser(email)

	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (s *ticketService) ViewSeatAllocation(ctx context.Context, section string) ([]*models.Seat, error) {
	seats, sErr := s.repo.GetSeatMapBySection(section)
	if sErr != nil {
		return nil, sErr
	}
	return seats, nil
}

func (s *ticketService) RemoveUserfromTrain(ctx context.Context, email string) error { panic("") }
func (s *ticketService) ModifyBooking(ctx context.Context, bookingID string) error   { panic("") }
