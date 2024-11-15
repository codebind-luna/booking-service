package inmemoryrepository

import (
	"errors"

	"github.com/codebind-luna/booking-service/internal/domain/models"
)

var (
	ErrSeatNotFound            = errors.New("seat not found")
	ErrSeatNotAvailableAnymore = errors.New("sorry requested seat is not available")
)

func (ir *InMemoryRepository) isValidSeat(seatNo int) bool {
	noOfSeatsEachSection := ir.seatM.Cols()
	seatNo -= 1
	return 0 <= seatNo && seatNo < noOfSeatsEachSection
}

func (ir *InMemoryRepository) ModifySeat(email string, section string, seatNo int) error {
	ir.mu.Lock()
	defer ir.mu.Unlock()

	s, parserErr := models.ParseSection(section)
	if parserErr != nil {
		return parserErr
	}

	u, exists := ir.users[email]
	if !exists {
		return ErrUserNotFound
	}

	t, ok := ir.userTicket[u.ID()]
	if !ok {
		return ErrNoBookingFoundForUser
	}

	seatIdx := seatNo - 1

	if !ir.isValidSeat(seatNo) {
		return ErrSeatNotFound
	}

	spot := ir.seatM.Seats()[s.EnumIndex()][seatIdx]

	if spot.Status() == models.Available {
		t.Seat().SetUser(nil)
		t.Seat().SetStatus(models.Available)

		spot.SetStatus(models.Booked)
		spot.SetUser(u)
		t.SetSeat(spot)
		return nil
	}

	return ErrSeatNotAvailableAnymore
}
