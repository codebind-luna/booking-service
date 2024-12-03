package ticket

import (
	"context"

	"github.com/codebind-luna/booking-service/internal/domain/models"
)

func (s *ticketService) ViewSeatMap(ctx context.Context, section string) ([]*models.Seat, error) {
	seats, sErr := s.repo.GetSeatMapBySection(section)
	if sErr != nil {
		return nil, sErr
	}
	return seats, nil
}
