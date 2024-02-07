package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"wblevel0/db"
	"wblevel0/server"
)

var (
	dbUrl        = flag.String("db_url", "postgres://wb:ok@localhost:5432/wb_level0?pool_max_conns=10", "The DB url connection string")
	waitShutdown = flag.Duration("waitShutdown", time.Second*15, "Duration wait seconds before shutdown server")
)

func main() {
	flag.Parse()

	database, err := db.NewDatabase(*dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	cacheManager := server.NewCacheManager()

	databus := server.NewDataBus(database, cacheManager)
	if err := databus.Run(); err != nil {
		log.Fatal(err)
	}
	defer databus.Stop()

	httpServer := server.NewHttpServer(cacheManager, 8080)
	go func() {
		httpServer.Run()
	}()

	err = restoreCache(database, cacheManager)
	if err != nil {
		log.Println(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), *waitShutdown)
	defer cancel()
	httpServer.Stop(ctx)
	log.Println("shutting down")

	os.Exit(0)
}

func restoreCache(database *db.WBDatabase, cacheManager *server.CacheManager) error {
	allOrders, err := database.OrderDao.GetAll()
	if err != nil {
		return fmt.Errorf("can not read orders from Database: %v", err)
	}
	for _, item := range allOrders {
		cacheManager.Add(item.OrderUID, item)
	}
	log.Printf("Resored cache with records: %v", len(allOrders))
	return nil
}
