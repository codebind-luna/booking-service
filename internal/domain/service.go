package domain

import (
	"context"

	"github.com/codebind-luna/booking-service/internal/domain/models"
)

type Service interface {
	PurchaseTicket(ctx context.Context, email, firstName, lastName, fromCity, toCity string, price float64) (*string, error)
	RemoveUserfromTrain(ctx context.Context, email string) error
	ModifySeat(ctx context.Context, email string, section string, seatNo int) error
	GetReceipt(ctx context.Context, bookingID string) (*models.Ticket, error)
	ViewSeatMap(ctx context.Context, section string) ([]*models.Seat, error)
}
