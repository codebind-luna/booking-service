package exceptions

import "errors"

var (
	ErrNoSeatsAvailable        = errors.New("sorry all seats are booked")
	ErrUserNotFound            = errors.New("user not found")
	ErrNoBookingFoundForUser   = errors.New("no tickets are for the user")
	ErrSeatNotFound            = errors.New("seat not found")
	ErrSeatNotAvailableAnymore = errors.New("sorry requested seat is not available")
)