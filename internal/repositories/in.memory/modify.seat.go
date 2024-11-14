package inmemoryrepository

import "github.com/codebind-luna/booking-service/internal/domain/models"

func (ir *InMemoryRepository) ModifySeat(user *models.User) error {
	ir.mu.Lock()
	defer ir.mu.Unlock()

	u, exists := ir.users[user.Email()]
	if !exists {
		return ErrUserNotFound
	}

	t, ok := ir.userTicket[u.ID()]
	if !ok {
		return ErrNoBookingFoundForUser
	}

	spot, spotErr := ir.allocateSeat()

	if spotErr != nil {
		return spotErr
	}

	t.Seat().SetUser(nil)
	t.Seat().SetStatus(models.Available)

	spot.SetUser(u)
	t.SetUser(u)
	t.SetSeat(spot)

	return nil
}
