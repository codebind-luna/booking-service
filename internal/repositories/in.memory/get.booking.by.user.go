package inmemoryrepository

import "github.com/codebind-luna/booking-service/internal/domain/models"

func (ir *InMemoryRepository) GetReceiptByUser(email string) (*models.Ticket, error) {
	ir.mu.RLock()
	defer ir.mu.RUnlock()

	u, exists := ir.users[email]
	if !exists {
		return nil, ErrUserNotFound
	}

	ticket, ok := ir.userTicket[u.ID()]
	if !ok {
		return nil, ErrNoBookingFoundForUser
	}

	return ticket, nil
}
