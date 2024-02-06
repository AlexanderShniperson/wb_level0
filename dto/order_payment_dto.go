package dto

import "wblevel0/db/entity"

type OrderPaymentDTO struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

func NewOrderPaymentDTOFromEntity(entity *entity.OrderEntity) *OrderPaymentDTO {
	return &OrderPaymentDTO{
		Transaction:  entity.Payment.Transaction,
		RequestId:    entity.Payment.RequestId,
		Currency:     entity.Payment.Currency,
		Provider:     entity.Payment.Provider,
		Amount:       entity.Payment.Amount,
		PaymentDt:    entity.Payment.PaymentDt,
		Bank:         entity.Payment.Bank,
		DeliveryCost: entity.Payment.DeliveryCost,
		GoodsTotal:   entity.Payment.GoodsTotal,
		CustomFee:    entity.Payment.CustomFee,
	}
}

func (e *OrderPaymentDTO) ToEntity() *entity.OrderPaymentEntity {
	return &entity.OrderPaymentEntity{
		Transaction:  e.Transaction,
		RequestId:    e.RequestId,
		Currency:     e.Currency,
		Provider:     e.Provider,
		Amount:       e.Amount,
		PaymentDt:    e.PaymentDt,
		Bank:         e.Bank,
		DeliveryCost: e.DeliveryCost,
		GoodsTotal:   e.GoodsTotal,
		CustomFee:    e.CustomFee,
	}
}
