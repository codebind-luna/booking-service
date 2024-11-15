package inmemoryrepository

import (
	"github.com/codebind-luna/booking-service/internal/domain/models"
	"github.com/codebind-luna/booking-service/internal/exceptions"
)

func (ir *InMemoryRepository) RemoveUser(email string) error {
	ir.mu.Lock()
	defer ir.mu.Unlock()

	u, exists := ir.users[email]
	if !exists {
		return exceptions.ErrNoBookingFoundForUser
	}

	t, ok := ir.userTicket[u.ID()]
	if !ok {
		return exceptions.ErrNoBookingFoundForUser
	}

	seat := t.Seat()
	seat.SetUser(nil)
	seat.SetStatus(models.Available)

	delete(ir.tickets, t.BookingID())
	delete(ir.userTicket, u.ID())

	return nil
}
