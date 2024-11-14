package inmemoryrepository

import "github.com/codebind-luna/booking-service/internal/domain/models"

func (ir *InMemoryRepository) addOrFetchUser(user *models.User) {
	_, exists := ir.users[user.Email()]
	if !exists {
		user.SetID()
		ir.users[user.Email()] = user
	}
}
