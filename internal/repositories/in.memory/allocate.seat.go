package inmemoryrepository

import (
	"time"

	"math/rand"

	"github.com/codebind-luna/booking-service/internal/domain/models"
)

func (ir *InMemoryRepository) allocateSeat() (*models.Seat, error) {
	// Find all empty spots
	emptySpots := ir.findEmptySpots()
	if len(emptySpots) == 0 {
		return nil, ErrNoSeatsAvailable
	}

	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Pick a random spot from the list of empty spots
	randomIndex := rand.Intn(len(emptySpots))
	spot := emptySpots[randomIndex]

	spot.SetStatus(models.Booked)

	return spot, nil
}
