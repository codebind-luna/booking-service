package ticket

import (
	"context"

	"github.com/codebind-luna/booking-service/internal/domain/models"
)

func (s *ticketService) GetReceipt(ctx context.Context, email string) (*models.Ticket, error) {
	ticket, err := s.repo.GetReceiptByUser(email)

	if err != nil {
		return nil, err
	}
	return ticket, nil
}
