package coupon

import (
	"errors"

	"github.com/codebind-luna/booking-service/internal/domain"
	"github.com/codebind-luna/booking-service/internal/domain/models"
)

// New - retrieve a coupon
func NewCoupon(
	couponCode models.CouponCode,
) (domain.Coupon, error) {
	switch couponCode {
	case models.DiscountFivePerCent:
		return NewFivePercentDiscountCoupon(models.DiscountFivePerCent), nil
	case models.DiscountTenPercent:
		return NewTenPercentDiscountCoupon(models.DiscountFivePerCent), nil
	default:
		return nil, errors.New("invalid coupon provided")
	}
}

// NewCoupons function to initialize and return a map of coupon codes to Coupons
func NewCoupons() map[string]domain.Coupon {
	// Initialize the map of coupons
	coupons := make(map[string]domain.Coupon)

	for k, v := range models.CouponCodeMap {
		coupon, _ := NewCoupon(v)
		coupons[k] = coupon
	}

	// Return the map of coupons
	return coupons
}
