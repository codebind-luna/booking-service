package inmemoryrepository

import (
	"github.com/codebind-luna/booking-service/internal/domain/models"
)

func (ir *InMemoryRepository) CreateBooking(user *models.User, fromCity string, toCity string, price float64) (*string, error) {
	ir.mu.Lock()
	defer ir.mu.Unlock()

	ir.addOrFetchUser(user)

	seat, allocationErr := ir.allocateSeat()
	if allocationErr != nil {
		return nil, allocationErr
	}

	seat.SetUser(user)

	booking := models.NewBooking(user, fromCity, toCity, seat, price)

	bookingID := booking.BookingID()

	ir.tickets[bookingID] = booking
	ir.userTicket[user.ID()] = booking
	return &bookingID, nil
}


