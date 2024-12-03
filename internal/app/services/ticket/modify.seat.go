package ticket

import "context"

func (s *ticketService) ModifySeat(ctx context.Context, email string, section string, seatNo int) error {
	err := s.repo.ModifySeat(email, section, seatNo)

	if err != nil {
		return err
	}

	return nil
}
