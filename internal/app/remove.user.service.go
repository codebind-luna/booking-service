package app

import "context"

func (s *ticketService) RemoveUserfromTrain(ctx context.Context, email string) error {
	err := s.repo.RemoveUser(email)

	if err != nil {
		return err
	}

	return nil
}
