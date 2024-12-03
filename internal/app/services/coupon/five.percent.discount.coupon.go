package coupon

import (
	"context"

	"github.com/codebind-luna/booking-service/internal/domain/models"
)

type fivePercentDiscountCoupon struct {
	code          models.CouponCode
	discount      float64
	isAlwaysValid bool
}

func NewFivePercentDiscountCoupon(code models.CouponCode) *fivePercentDiscountCoupon {
	return &fivePercentDiscountCoupon{
		isAlwaysValid: true,
		discount:      5,
		code:          code,
	}
}

func (fc *fivePercentDiscountCoupon) Valid(ctx context.Context) bool {
	return fc.isAlwaysValid
}

func (fc *fivePercentDiscountCoupon) ApplyCoupon(ctx context.Context, price float64) float64 {
	return price * (1 - fc.discount/100)
}
