package entity

import "time"

type OrderEntity struct {
	OrderUID          string
	TrackNumber       string
	Entry             string
	Delivery          *OrderDeliveryEntity
	Payment           *OrderPaymentEntity
	Items             []*OrderItemEntity
	Locale            string
	InternalSignature string
	CustomerId        string
	DeliveryService   string
	ShardKey          string
	SmId              int
	DateCreated       time.Time
	OofShard          string
}
