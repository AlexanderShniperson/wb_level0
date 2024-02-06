package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"wblevel0/db"
	"wblevel0/db/entity"
	"wblevel0/dto"

	"github.com/gorilla/mux"
	nats "github.com/nats-io/nats.go"
)

var (
	dbUrl     = flag.String("db_url", "postgres://wb:ok@localhost:5432/wb_level0?pool_max_conns=10", "The DB url connection string")
	dataCache = make(map[string]*entity.OrderEntity)
)

func main() {
	flag.Parse()

	database, err := db.NewDatabase(*dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	sc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("databus", func(m *nats.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
		orderDto := &dto.OrderDTO{}
		err := json.Unmarshal(m.Data, &orderDto)
		if err == nil {
			log.Printf("Parsed OrderDTO with ID: %s\n", orderDto.OrderUID)
			orderEntity := orderDto.ToEntity()
			err := database.OrderDao.Add(orderEntity)
			if err == nil {
				log.Printf("Added new Order with ID: %s\n", orderDto.OrderUID)
				dataCache[orderEntity.OrderUID] = orderEntity
			} else {
				log.Printf("Error: %v\n", err)
			}
		} else {
			log.Printf("Error: %v\n", err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePageHandler)
	myRouter.HandleFunc("/api/order/{id}", orderDetailHandler)
	go func() {
		http.ListenAndServe(":8080", myRouter)
	}()

	allOrders, err := database.OrderDao.GetAll()
	if err != nil {
		log.Fatalf("Can not read orders from Database: %v", err)
	}
	for i := len(allOrders) - 1; i >= 0; i-- {
		item := allOrders[i]
		dataCache[item.OrderUID] = item
		allOrders = allOrders[:len(allOrders)-1]
	}

	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Press ENTER to exit.")
	if scanner.Scan() {
		log.Println("Exit from application, goodbye.")
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/index.html")
}

func orderDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)

	orderId, paramExists := vars["id"]

	if !paramExists {
		http.Error(w, "Error: wrong parameter orderId.", http.StatusBadRequest)
		return
	}

	order, orderExists := dataCache[orderId]
	if !orderExists {
		http.Error(w, "Error: order does not exist.", http.StatusBadRequest)
		return
	}

	orderDto := dto.NewOrderDTOFromEntity(order)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderDto)
}
