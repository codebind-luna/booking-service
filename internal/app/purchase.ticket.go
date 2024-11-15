package app

import (
	"context"

	"github.com/codebind-luna/booking-service/internal/domain/models"
)

func (s *ticketService) PurchaseTicket(ctx context.Context, email, firstName, lastName, fromCity, toCity string, price float64) (*string, error) {
	user := models.NewUser(email, firstName, lastName)

	bookingID, bookingErr := s.repo.CreateBooking(user, fromCity, toCity, price)

	if bookingErr != nil {
		return nil, bookingErr
	}

	return bookingID, nil
}
