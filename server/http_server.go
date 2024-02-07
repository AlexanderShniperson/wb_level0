package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"wblevel0/dto"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	cacheManager *CacheManager
	srv          *http.Server
}

func NewHttpServer(cacheManager *CacheManager, port int) *HttpServer {
	router := mux.NewRouter().StrictSlash(true)
	result := &HttpServer{
		cacheManager: cacheManager,
		srv: &http.Server{
			Handler: router,
			Addr:    fmt.Sprintf(":%v", port),
		},
	}
	router.HandleFunc("/", result.homePageHandler)
	router.HandleFunc("/api/order/{id}", result.orderDetailHandler)
	return result
}

func (e *HttpServer) Run() error {
	return e.srv.ListenAndServe()
}

func (e *HttpServer) Stop(context context.Context) error {
	return e.srv.Shutdown(context)
}

func (e *HttpServer) homePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page"))
}

func (e *HttpServer) orderDetailHandler(w http.ResponseWriter, r *http.Request) {
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

	orderEntity, err := e.cacheManager.Get(orderId)
	if err != nil {
		http.Error(w, "Error: order does not exist.", http.StatusBadRequest)
		return
	}

	orderDto := dto.NewOrderDTOFromEntity(orderEntity)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderDto)
}
