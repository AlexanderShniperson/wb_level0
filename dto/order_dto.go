package dto

import (
	"time"
	"wblevel0/db/entity"
)

type OrderDTO struct {
	OrderUID          string            `json:"order_uid"`
	TrackNumber       string            `json:"track_number"`
	Entry             string            `json:"entry"`
	Delivery          *OrderDeliveryDTO `json:"delivery"`
	Payment           *OrderPaymentDTO  `json:"payment"`
	Items             []*OrderItemDTO   `json:"items"`
	Locale            string            `json:"locale"`
	InternalSignature string            `json:"internal_signature"`
	CustomerId        string            `json:"customer_id"`
	DeliveryService   string            `json:"delivery_service"`
	ShardKey          string            `json:"shardkey"`
	SmId              int               `json:"sm_id"`
	DateCreated       time.Time         `json:"date_created"`
	OofShard          string            `json:"oof_shard"`
}

func NewOrderDTOFromEntity(entity *entity.OrderEntity) *OrderDTO {
	items := make([]*OrderItemDTO, len(entity.Items))
	for idx, orderItem := range entity.Items {
		items[idx] = NewOrderItemDTOFromEntity(orderItem)
	}
	return &OrderDTO{
		OrderUID:          entity.OrderUID,
		TrackNumber:       entity.TrackNumber,
		Entry:             entity.Entry,
		Delivery:          NewOrderDeliveryDTOFromEntity(entity),
		Payment:           NewOrderPaymentDTOFromEntity(entity),
		Items:             items,
		Locale:            entity.Locale,
		InternalSignature: entity.InternalSignature,
		CustomerId:        entity.CustomerId,
		DeliveryService:   entity.DeliveryService,
		ShardKey:          entity.ShardKey,
		SmId:              entity.SmId,
		OofShard:          entity.OofShard,
		DateCreated:       entity.DateCreated,
	}
}

func (e *OrderDTO) ToEntity() *entity.OrderEntity {
	items := make([]*entity.OrderItemEntity, len(e.Items))
	for idx, orderItem := range e.Items {
		items[idx] = orderItem.ToEntity()
	}
	return &entity.OrderEntity{
		OrderUID:          e.OrderUID,
		TrackNumber:       e.TrackNumber,
		Entry:             e.Entry,
		Delivery:          e.Delivery.ToEntity(),
		Payment:           e.Payment.ToEntity(),
		Items:             items,
		Locale:            e.Locale,
		InternalSignature: e.InternalSignature,
		CustomerId:        e.CustomerId,
		DeliveryService:   e.DeliveryService,
		ShardKey:          e.ShardKey,
		SmId:              e.SmId,
		DateCreated:       e.DateCreated,
		OofShard:          e.OofShard,
	}
}
