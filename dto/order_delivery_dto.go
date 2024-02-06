package dto

import "wblevel0/db/entity"

type OrderDeliveryDTO struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

func NewOrderDeliveryDTOFromEntity(entity *entity.OrderEntity) *OrderDeliveryDTO {
	return &OrderDeliveryDTO{
		Name:    entity.Delivery.Name,
		Phone:   entity.Delivery.Phone,
		Zip:     entity.Delivery.Zip,
		City:    entity.Delivery.City,
		Address: entity.Delivery.Address,
		Region:  entity.Delivery.Region,
		Email:   entity.Delivery.Email,
	}
}

func (e *OrderDeliveryDTO) ToEntity() *entity.OrderDeliveryEntity {
	return &entity.OrderDeliveryEntity{
		Name:    e.Name,
		Phone:   e.Phone,
		Zip:     e.Zip,
		City:    e.City,
		Address: e.Address,
		Region:  e.Region,
		Email:   e.Email,
	}
}
