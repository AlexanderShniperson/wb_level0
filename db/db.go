package db

import (
	"context"
	"fmt"
	"wblevel0/db/dao"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WBDatabase struct {
	dbPool   *pgxpool.Pool
	OrderDao *dao.OrderDao
}

func NewDatabase(dbUrl string) (*WBDatabase, error) {
	dbconfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pool config: %v", err)
	}
	dbPool, err := pgxpool.NewWithConfig(context.Background(), dbconfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}
	return &WBDatabase{
		dbPool:   dbPool,
		OrderDao: dao.NewOrderDao(dbPool),
	}, nil
}

func (e *WBDatabase) Close() {
	e.dbPool.Close()
}
