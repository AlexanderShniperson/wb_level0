package server

import (
	"errors"
	"wblevel0/db/entity"
)

type CacheManager struct {
	dataCache map[string]*entity.OrderEntity
}

func NewCacheManager() *CacheManager {
	return &CacheManager{
		dataCache: make(map[string]*entity.OrderEntity),
	}
}

func (e *CacheManager) Get(key string) (*entity.OrderEntity, error) {
	if item, ok := e.dataCache[key]; ok {
		return item, nil
	}
	return nil, errors.New("item not found")
}

func (e *CacheManager) Add(key string, value *entity.OrderEntity) {
	e.dataCache[key] = value
}
