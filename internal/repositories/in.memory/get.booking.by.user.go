package inmemoryrepository

import (
	"github.com/codebind-luna/booking-service/internal/domain/models"
	"github.com/codebind-luna/booking-service/internal/exceptions"
)

func (ir *InMemoryRepository) GetReceiptByUser(email string) (*models.Ticket, error) {
	ir.mu.RLock()
	defer ir.mu.RUnlock()

	u, exists := ir.users[email]
	if !exists {
		return nil, exceptions.ErrNoBookingFoundForUser
	}

	ticket, ok := ir.userTicket[u.ID()]
	if !ok {
		return nil, exceptions.ErrNoBookingFoundForUser
	}

	return ticket, nil
}
