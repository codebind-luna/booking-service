package inmemoryrepository

import "github.com/codebind-luna/booking-service/internal/domain/models"

func (ir *InMemoryRepository) GetSeatMapBySection(section string) ([]*models.Seat, error) {
	s, parserErr := models.ParseSection(section)
	if parserErr != nil {
		return nil, parserErr
	}

	ir.mu.RLock()
	defer ir.mu.RUnlock()

	return ir.seatM.Seats()[s.EnumIndex()], nil
}
