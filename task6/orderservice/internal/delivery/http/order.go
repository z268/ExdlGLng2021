package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	pb "github.com/z268/ExdlGLng2021/task6/orderservice/internal/delivery/grpc"
	"github.com/z268/ExdlGLng2021/task6/orderservice/internal/repository"
)

func (h *OrderHandler) initOrderRoutes(r *mux.Router) {
	listRoute := r.PathPrefix("/orders").Subrouter()
	listRoute.HandleFunc("/", h.orderList).Methods("GET")
	listRoute.HandleFunc("/", h.orderCreate).Methods("POST")

	detailRoute := listRoute.PathPrefix("/{order_uuid:[a-f0-9-]+}").Subrouter()
	detailRoute.HandleFunc("/", h.orderDetail).Methods("GET")
	detailRoute.HandleFunc("/", h.orderUpdate).Methods("PUT", "PATCH")
	detailRoute.HandleFunc("/", h.orderDelete).Methods("DELETE")
}

func (h *OrderHandler) getOrderFromRequest(r *http.Request) (*repository.Order, error) {
	id, err := uuid.Parse(mux.Vars(r)["order_uuid"])
	if err != nil {
		return nil, err
	}

	return h.repo.Get(id)
}

// ListOrder godoc
// @Summary Create new order
// @Tags Orders
// @Param message body CreateOrderRequest true "Order"
// @Success 201 {object} repository.Order
// @Failure 500 {object} ResponseError
// @Router /orders/ [post]
func (h *OrderHandler) orderCreate(w http.ResponseWriter, r *http.Request) {
	// TODO: add input validation

	order := &repository.Order{UUID: uuid.New()}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	grpcRequest := &pb.GetBooksRequest{BookUuid: []string{order.Book_uuid.String()}}
	_, err = h.grpcClient.GetBookByUUID(ctx, grpcRequest)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	order, err = h.repo.Create(order)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, order, http.StatusOK)
}

// CreateOrder godoc
// @Summary Get all orders
// @Tags Orders
// @Success 200 {array} repository.Order
// @Router /orders/ [get]
func (h *OrderHandler) orderList(w http.ResponseWriter, r *http.Request) {
	orders, err := h.repo.List()
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, orders, http.StatusOK)
}

// RetriveOrder godoc
// @Summary Get order detail
// @Tags Orders
// @Success 200 {object} repository.Order
// @Failure 404 {object} ResponseError
// @Param order_uuid path string true "Order UUID"
// @Router /orders/{order_uuid}/ [get]
func (h *OrderHandler) orderDetail(w http.ResponseWriter, r *http.Request) {
	order, err := h.getOrderFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	SendResponse(w, order, http.StatusOK)
}

// UpdateOrder godoc
// @Summary Update an order
// @Tags Orders
// @Success 200 {object} repository.Order
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Param order_uuid path string true "Order UUID"
// @Param message body UpdateOrderRequest true "Order"
// @Router /orders/{order_uuid}/ [put]
func (h *OrderHandler) orderUpdate(w http.ResponseWriter, r *http.Request) {
	// TODO: add input validation for being sure that we update only description field

	order, err := h.getOrderFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&order); err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	if err = h.repo.Update(order); err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, order, http.StatusOK)
}

// DeleteOrder godoc
// @Summary Delete an order
// @Tags Orders
// @Success 200 {string} string	"Empty response"
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Param order_uuid path string true "Order UUID"
// @Router /orders/{order_uuid}/ [delete]
func (h *OrderHandler) orderDelete(w http.ResponseWriter, r *http.Request) {
	order, err := h.getOrderFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	if err = h.repo.Delete(order); err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
