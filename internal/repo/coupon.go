package repo

import (
	"context"
	"github.com/endr-i/promo-back-go/internal/model"
	"github.com/jmoiron/sqlx"
)

type CouponDto struct {
	Code         string
	CouponTypeId string
}

type CouponRepo interface {
	Create(ctx context.Context, dto CouponDto) (*model.Coupon, error)
}

type couponRepo struct {
	db sqlx.ExtContext
}

func (r *couponRepo) Create(ctx context.Context, dto CouponDto) (*model.Coupon, error) {
	var coupon model.Coupon
	err := r.db.QueryRowxContext(
		ctx,
		"INSERT INTO coupons (code, coupon_type_id) VALUES (?, ?) RETURNING *",
		dto.Code,
		dto.CouponTypeId,
	).StructScan(&coupon)
	if err != nil {
		return nil, err
	}

	return &coupon, nil
}

func NewCouponRepo(db sqlx.ExtContext) CouponRepo {
	return &couponRepo{db: db}
}
