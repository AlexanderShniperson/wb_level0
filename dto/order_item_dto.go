package dto

import "wblevel0/db/entity"

type OrderItemDTO struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	RId         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func NewOrderItemDTOFromEntity(entity *entity.OrderItemEntity) *OrderItemDTO {
	return &OrderItemDTO{
		ChrtId:      entity.ChrtId,
		TrackNumber: entity.TrackNumber,
		Price:       entity.Price,
		RId:         entity.RId,
		Name:        entity.Name,
		Sale:        entity.Sale,
		Size:        entity.Size,
		TotalPrice:  entity.TotalPrice,
		NmId:        entity.NmId,
		Brand:       entity.Brand,
		Status:      entity.Status,
	}
}

func (e *OrderItemDTO) ToEntity() *entity.OrderItemEntity {
	return &entity.OrderItemEntity{
		ChrtId:      e.ChrtId,
		TrackNumber: e.TrackNumber,
		Price:       e.Price,
		RId:         e.RId,
		Name:        e.Name,
		Sale:        e.Sale,
		Size:        e.Size,
		TotalPrice:  e.TotalPrice,
		NmId:        e.NmId,
		Brand:       e.Brand,
		Status:      e.Status,
	}
}
