package http

import (
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"

	_"github.com/z268/ExdlGLng2021/task6/catalogservice/docs"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
)

type CatalogHandler struct {
	repo          *repository.CatalogRepository
	apiPageSize   uint64
}

func NewCatalogHandler(repo *repository.CatalogRepository, apiPageSize uint64) *CatalogHandler {
	return &CatalogHandler{repo, apiPageSize}
}

func InitCatalogRouter(handler *CatalogHandler) *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	apiV1 := r.PathPrefix("/api/v1").Subrouter()
	handler.Init(apiV1)

	return r
}

func (h *CatalogHandler) Init(r *mux.Router) {
	h.initBookRoutes(r)
	h.initAuthorRoutes(r)
	h.initCategoryRoutes(r)
}
