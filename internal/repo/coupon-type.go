package repo

import (
	"context"
	"github.com/endr-i/promo-back-go/internal/model"
	"github.com/jmoiron/sqlx"
)

type CouponTypeDto struct {
	Name   string `json:"name"`
	Chance int    `json:"chance"`
}

type CouponTypeRepo interface {
	Create(ctx context.Context, row CouponTypeDto) (*model.CouponType, error)
}

type couponTypeRepo struct {
	db sqlx.ExtContext
}

func (r *couponTypeRepo) Create(ctx context.Context, row CouponTypeDto) (*model.CouponType, error) {
	var couponType model.CouponType
	err := r.db.QueryRowxContext(
		ctx,
		"INSERT INTO coupon_types (name, chance) VALUES ($1, $2) RETURNING *",
		row.Name,
		row.Chance,
	).StructScan(&couponType)
	if err != nil {
		return nil, err
	}

	return &couponType, nil
}

func NewCouponTypeRepo(db sqlx.ExtContext) CouponTypeRepo {
	return &couponTypeRepo{db: db}
}
