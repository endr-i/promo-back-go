package model

type Coupon struct {
	Id           string `db:"id"`
	Code         string `db:"code"`
	CouponTypeId string `db:"coupon_type_id"`
}
