package server

import (
	"encoding/json"
	"log"
	"wblevel0/db"
	"wblevel0/dto"

	nats "github.com/nats-io/nats.go"
)

type DataBus struct {
	database     *db.WBDatabase
	cacheManager *CacheManager
	conn         *nats.Conn
	subs         *nats.Subscription
}

func NewDataBus(database *db.WBDatabase, cacheManager *CacheManager) *DataBus {
	return &DataBus{
		database:     database,
		cacheManager: cacheManager,
	}
}

func (e *DataBus) Run() error {
	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	natsSubs, err := natsConn.Subscribe("databus", func(m *nats.Msg) {
		orderDto := &dto.OrderDTO{}
		err := json.Unmarshal(m.Data, &orderDto)
		if err == nil {
			log.Printf("Parsed OrderDTO with ID: %s\n", orderDto.OrderUID)
			orderEntity := orderDto.ToEntity()
			err := e.database.OrderDao.Add(orderEntity)
			if err == nil {
				log.Printf("Added new Order with ID: %s\n", orderDto.OrderUID)
				e.cacheManager.Add(orderEntity.OrderUID, orderEntity)
			} else {
				log.Println("Error add Order into database")
			}
		} else {
			log.Println("Error parse OrderDTO")
		}
	})
	if err != nil {
		natsConn.Close()
		return err
	}
	e.conn = natsConn
	e.subs = natsSubs
	return nil
}

func (e *DataBus) Stop() error {
	err := e.subs.Unsubscribe()
	e.conn.Close()
	return err
}
