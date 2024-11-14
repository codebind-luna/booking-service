package inmemoryrepository

import "github.com/codebind-luna/booking-service/internal/domain/models"

func (ir *InMemoryRepository) findEmptySpots() []*models.Seat {
	var emptySpots []*models.Seat
	seats := ir.seatM.Seats()
	for i := 0; i < 2; i++ {
		for j := 0; j < ir.seatM.Cols(); j++ {
			if seats[i][j].IsAvailable() {
				emptySpots = append(emptySpots, seats[i][j])
			}
		}
	}
	return emptySpots
}
