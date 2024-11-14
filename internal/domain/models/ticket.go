package models

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	bookingID  string
	user       *User
	fromCity   string
	toCity     string
	price      float64
	bookedAt   int64
	seat       *Seat
	modifiedAt int64
}

func NewBooking(user *User, from, to string, seat *Seat, price float64) *Ticket {
	id := uuid.NewString()
	return &Ticket{
		bookingID: id,
		user:      user,
		fromCity:  from,
		toCity:    to,
		seat:      seat,
		price:     price,
		bookedAt:  time.Now().Unix(),
	}
}

func (b *Ticket) User() *User {
	return b.user
}

func (b *Ticket) BookingID() string {
	return b.bookingID
}

func (b *Ticket) FromCity() string {
	return b.fromCity
}

func (b *Ticket) ToCity() string {
	return b.toCity
}

func (b *Ticket) PricePaid() float64 {
	return b.price
}

func (b *Ticket) Seat() *Seat {
	return b.seat
}

func (b *Ticket) SetSeat(s *Seat) {
	b.seat = s
}

func (b *Ticket) SetUser(u *User) {
	b.user = u
}

func (b *Ticket) SetModifiedAt() {
	b.modifiedAt = time.Now().Unix()
}
