package handlers

import (
	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
	"github.com/codebind-luna/booking-service/internal/app/services/coupon"
	"github.com/codebind-luna/booking-service/internal/domain"
)

type ticketServiceHandlers struct {
	bookingv1.UnimplementedTicketServiceServer
	svc     domain.Service
	coupons map[string]domain.Coupon
}

func NewTicketService(svc domain.Service) *ticketServiceHandlers {
	coupons := coupon.NewCoupons()
	return &ticketServiceHandlers{
		coupons: coupons,
		svc:     svc,
	}
}
