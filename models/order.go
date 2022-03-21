package models

type OrderDTO struct {
	OrderId           int    `db:"id" json:"-"`
	OrderUID          string `db:"order_uid" json:"order_uid" validate:"required"`
	TrackNumber       string `db:"track_number" json:"track_number" validate:"required"`
	Entry             string `db:"entry" json:"entry" validate:"required"`
	Local             string `db:"local" json:"locale" validate:"required"`
	InternalSignature string `db:"internal_signature" json:"internal_signature" validate:"omitempty"`
	CustomerID        string `db:"customer_id" json:"customer_id" validate:"required"`
	DeliveryService   string `db:"delivery_service" json:"delivery_service" validate:"required"`
	Shardkey          string `db:"shardkey" json:"shardkey" validate:"required"`
	SmID              int64  `db:"sm_id" json:"sm_id" validate:"required"`
	DateCreated       string `db:"date_created" json:"date_created" validate:"required"`
	OofShard          string `db:"oof_shard" json:"oof_shard" validate:"required"`

	DeliveryId int         `db:"delivery_id" json:"-"`
	Delivery   DeliveryDTO `db:"delivery" json:"delivery"`
	PaymentId  int         `db:"payment_id" json:"-"`
	Payment    PaymentDTO  `db:"payment" json:"payment"`
	Items      []ItemDTO   `db:"items" validate:"required,dive,required" json:"items"`
}

type PaymentDTO struct {
	Id           int    `db:"id" json:"-"`
	Transaction  string `db:"transaction" json:"transaction" validate:"required"`
	RequestID    string `db:"request_id" json:"request_id" validate:"required"`
	Currency     string `db:"currency" json:"currency" validate:"required"`
	Provider     string `db:"provider" json:"provider" validate:"required"`
	Amount       int64  `db:"amount" json:"amount" validate:"required"`
	PaymentDt    int64  `db:"payment_dt" json:"payment_dt" validate:"required"`
	Bank         string `db:"bank" json:"bank" validate:"required"`
	DeliveryCost int64  `db:"delivery_cost" json:"delivery_cost" validate:"required"`
	GoodsTotal   int64  `db:"goods_total" json:"goods_total" validate:"required"`
	CustomFee    int64  `db:"custom_fee" json:"custom_fee" validate:"required"`
}

type ItemDTO struct {
	Id          int    `db:"id" json:"-"`
	ChrtId      int64  `db:"chrt_id" json:"chrt_id" validate:"required"`
	TrackNumber string `db:"track_number" json:"track_number" validate:"required,max=256"`
	Price       int64  `db:"price" json:"price" validate:"required"`
	Rid         string `db:"rid" json:"rid" validate:"required"`
	Name        string `db:"name" json:"name" validate:"required,max=128"`
	Sale        int64  `db:"sale" json:"sale" validate:"required"`
	Size        string `db:"size" json:"size" validate:"required"`
	TotalPrice  int64  `db:"total_price" json:"total_price" validate:"required"`
	NmID        int64  `db:"nm_id" json:"nm_id" validate:"required"`
	Brand       string `db:"brand" json:"brand" validate:"required,max=256"`
	Status      int64  `db:"status" json:"status" validate:"required"`
}

type DeliveryDTO struct {
	Id      int    `db:"id" json:"-"`
	Name    string `db:"name" json:"name" validate:"required"`
	Phone   string `db:"phone" json:"phone" validate:"required"`
	Zip     string `db:"zip" json:"zip" validate:"required"`
	City    string `db:"city" json:"city" validate:"required"`
	Address string `db:"address" json:"address" validate:"required"`
	Region  string `db:"region" json:"region" validate:"required"`
	Email   string `db:"email" json:"email" validate:"required"`
}

type ItemsOrdersDTO struct {
	OrderId int `db:"order_id"`
	ItemId  int `db:"item_id"`
}
