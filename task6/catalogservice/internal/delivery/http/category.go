package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
	"net/http"
)

func (h *CatalogHandler) initCategoryRoutes(r *mux.Router) {
	listRouter := r.PathPrefix("/categories").Subrouter()
	listRouter.HandleFunc("/", h.categoryList).Methods("GET")
	listRouter.HandleFunc("/", h.categoryCreate).Methods("POST")

	detailRouter := listRouter.PathPrefix("/{category_uuid:[a-f0-9-]+}").Subrouter()
	detailRouter.HandleFunc("/", h.categoryDetail).Methods("GET")
	detailRouter.HandleFunc("/", h.categoryUpdate).Methods("PUT", "PATCH")
	detailRouter.HandleFunc("/", h.categoryDelete).Methods("DELETE")
}

type CategoryRequest struct {
	Name        string    `json:"name"        example:"Programming"`
	Parent_uuid uuid.UUID `json:"parent_uuid" example:"524ab1ae-3293-4e90-9458-88b61eb95f0d"`
}

// @Summary Create an category
// @Tags Categories
// @Param message body CategoryRequest true "Category"
// @Success 201 {object} repository.Category
// @Failure 500 {object} ResponseError
// @Router /categories/ [post]
func (h *CatalogHandler) categoryCreate(w http.ResponseWriter, r *http.Request) {
	category := &repository.Category{UUID: uuid.New()}
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	category, err = h.repo.Category.Create(category)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, category, http.StatusCreated)
}

// @Summary Get list of categories
// @Tags Categories
// @Success 200 {array} repository.Category
// @Failure 500 {object} ResponseError
// @Router /categories/ [get]
func (h *CatalogHandler) categoryList(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.Category.List()
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, categories, http.StatusOK)
}

// @Summary Get a category
// @Tags Categories
// @Success 200 {object} repository.Category
// @Failure 404 {object} ResponseError
// @Param category_uuid path string true "Category UUID"
// @Router /categories/{category_uuid}/ [get]
func (h *CatalogHandler) categoryDetail(w http.ResponseWriter, r *http.Request) {
	category, err := h.getCategoryFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	SendResponse(w, category, http.StatusOK)
}


// @Summary Update a category
// @Tags Categories
// @Param category_uuid path string true "Category UUID"
// @Param message body CategoryRequest true "Category"
// @Success 200 {object} repository.Category
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /categories/{category_uuid}/ [put]
func (h *CatalogHandler) categoryUpdate(w http.ResponseWriter, r *http.Request) {
	category, err := h.getCategoryFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = h.repo.Category.Update(category)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, category, http.StatusOK)
}

// @Summary Delete a category
// @Tags Categories
// @Success 200 {string} string "Empty response"
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Param category_uuid path string true "Category UUID"
// @Router /categories/{category_uuid}/ [delete]
func (h *CatalogHandler) categoryDelete(w http.ResponseWriter, r *http.Request) {
	category, err := h.getCategoryFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	err = h.repo.Category.Delete(category)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}