package handlers

import (
	"context"
	"errors"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
	"github.com/codebind-luna/booking-service/internal/constants"
	"github.com/codebind-luna/booking-service/internal/exceptions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *ticketServiceHandlers) PurchaseTicket(ctx context.Context, pr *bookingv1.PurchaseTicketRequest) (*bookingv1.PurchaseTicketResponse, error) {
	user := pr.GetUser()
	fromCity := pr.GetFromCity()
	toCity := pr.GetToCity()
	price := pr.GetPrice()
	discount := pr.GetDiscount()

	discountedPrice, disErr := h.applyCoupon(ctx, float64(price), discount.String())
	if disErr != nil {
		return nil, status.New(codes.InvalidArgument, disErr.Error()).Err()
	}

	bookingID, err := h.svc.PurchaseTicket(ctx,
		user.GetEmail(), user.GetFirstName(), user.GetLastName(),
		fromCity, toCity, *discountedPrice)

	if err != nil {
		if errors.Is(err, exceptions.ErrNoSeatsAvailable) {
			return nil, status.New(codes.FailedPrecondition, "No seats available").Err()
		}
		if errors.Is(err, exceptions.ErrUserHasPurchasedTicketAlready) {
			return nil, status.New(codes.ResourceExhausted, "Already purchased a ticket").Err()
		}

		return nil, status.New(codes.Unknown, "some unknow error happened").Err()
	}

	r := &bookingv1.PurchaseTicketResponse{}
	r.BookingId = *bookingID
	r.Message = constants.SuccessPurchaseTicketMessage
	r.Success = true

	return r, nil
}

func (h *ticketServiceHandlers) applyCoupon(ctx context.Context, price float64, code string) (*float64, error) {
	if code == bookingv1.Discount_DISCOUNT_UNSPECIFIED.String() {
		return &price, nil
	}

	coupon, ok := h.coupons[code]

	if !ok {
		return nil, exceptions.ErrInvalidCouponCode
	}

	var discountedPrice float64

	if coupon.Valid(ctx) {
		discountedPrice = coupon.ApplyCoupon(ctx, float64(price))
	} else {
		return nil, exceptions.ErrCouponCodeExpired
	}
	return &discountedPrice, nil
}
