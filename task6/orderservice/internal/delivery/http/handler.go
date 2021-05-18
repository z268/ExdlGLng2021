package http

import (
	"github.com/gorilla/mux"

	"github.com/swaggo/http-swagger"
	_"github.com/z268/ExdlGLng2021/task6/orderservice/docs"

	"github.com/z268/ExdlGLng2021/task6/orderservice/internal/repository"
	"github.com/z268/ExdlGLng2021/task6/orderservice/internal/delivery/grpc"
)

type OrderHandler struct {
	repo repository.OrderRepository
	grpcClient grpc.BookServiceClient
}

func NewOrderHandler(repo repository.OrderRepository, grpcClient grpc.BookServiceClient) *OrderHandler {
	return &OrderHandler{repo, grpcClient}
}

func (h *OrderHandler) Init(r *mux.Router) {
	h.initOrderRoutes(r)
}

func InitOrderRouter(repo repository.OrderRepository, grpc_client grpc.BookServiceClient) *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	orderHandler := NewOrderHandler(repo, grpc_client)
	orderHandler.Init(apiV1)

	return r
}
