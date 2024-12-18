package inmemoryrepository

import (
	"os"
	"strconv"
	"sync"

	"github.com/codebind-luna/booking-service/internal/constants"
	"github.com/codebind-luna/booking-service/internal/domain/models"
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
func NewInMemoryRepository() *InMemoryRepository {
	noEachSection := os.Getenv(constants.EnvInmemorySeats)
	var seats int
	if noEachSection == "" {
		seats = constants.DefaultSeats
	} else {
		seats, _ = strconv.Atoi(noEachSection)
	}

	seatMap := models.NewSeatMap(seats)
	return &InMemoryRepository{
		tickets:    make(map[string]*models.Ticket),
		seatM:      seatMap,
		users:      make(map[string]*models.User),
		userTicket: make(map[string]*models.Ticket),
	}
}
