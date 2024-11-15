package inmemoryrepository

import "github.com/codebind-luna/booking-service/internal/domain/models"

func (ir *InMemoryRepository) RemoveUser(email string) error {
	ir.mu.Lock()
	defer ir.mu.Unlock()

	u, exists := ir.users[email]
	if !exists {
		return ErrUserNotFound
	}

	t, ok := ir.userTicket[u.ID()]
	if !ok {
		return ErrNoBookingFoundForUser
	}

	seat := t.Seat()
	seat.SetUser(nil)
	seat.SetStatus(models.Available)

	delete(ir.tickets, t.BookingID())
	delete(ir.userTicket, u.ID())

	return nil
}
