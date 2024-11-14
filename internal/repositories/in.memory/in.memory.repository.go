package inmemoryrepository

import (
	"errors"
	"sync"

	"github.com/codebind-luna/booking-service/internal/domain/models"
)

var (
	ErrNoSeatsAvailable      = errors.New("sorry all seats are booked")
	ErrUserNotFound          = errors.New("user not found")
	ErrNoBookingFoundForUser = errors.New("no tickets are for the user")
)

// InMemoryRepository is an in-memory implementation of the Repository interface.
type InMemoryRepository struct {
	users      map[string]*models.User
	tickets    map[string]*models.Ticket // Stores tickets by BookingID
	userTicket map[string]*models.Ticket
	seatM      models.SeatMap
	mu         sync.RWMutex
}

// NewInMemoryTicketRepository creates a new instance of InMemoryTicketRepository.
func NewInMemoryRepository(noEachSection int) *InMemoryRepository {
	seatMap := models.NewSeatMap(noEachSection)
	return &InMemoryRepository{
		tickets:    make(map[string]*models.Ticket),
		seatM:      seatMap,
		users:      make(map[string]*models.User),
		userTicket: make(map[string]*models.Ticket),
	}
}
