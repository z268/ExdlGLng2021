package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
	"net/http"
)

func (h *CatalogHandler) initAuthorRoutes(r *mux.Router) {
	listRoute := r.PathPrefix("/authors").Subrouter()
	listRoute.HandleFunc("/", h.authorList).Methods("GET")
	listRoute.HandleFunc("/", h.authorCreate).Methods("POST")

	detailRoute := listRoute.PathPrefix("/{author_uuid:[a-f0-9-]+}").Subrouter()
	detailRoute.HandleFunc("/", h.authorDetail).Methods("GET")
	detailRoute.HandleFunc("/", h.authorUpdate).Methods("PUT", "PATCH")
	detailRoute.HandleFunc("/", h.authorDelete).Methods("DELETE")
}

type AuthorRequest struct {
	Name string         `json:"name" example:"Robert Martin"`
}

// @Summary Create an author
// @Tags Authors
// @Param message body AuthorRequest true "Author"
// @Success 201 {object} repository.Author
// @Failure 500 {object} ResponseError
// @Router /authors/ [post]
func (h *CatalogHandler) authorCreate(w http.ResponseWriter, r *http.Request) {
	author := &repository.Author{UUID: uuid.New()}
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	author, err = h.repo.Author.Create(author)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, author, http.StatusCreated)
}

// @Summary Get list of authors
// @Tags Authors
// @Success 200 {array} repository.Author
// @Failure 500 {object} ResponseError
// @Router /authors/ [get]
func (h *CatalogHandler) authorList(w http.ResponseWriter, r *http.Request) {
	authors, err := h.repo.Author.List()
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, authors, http.StatusOK)
}

// @Summary Get an author
// @Tags Authors
// @Success 200 {object} repository.Author
// @Failure 404 {object} ResponseError
// @Param author_uuid path string true "Author UUID"
// @Router /authors/{author_uuid}/ [get]
func (h *CatalogHandler) authorDetail(w http.ResponseWriter, r *http.Request) {
	author, err := h.getAuthorFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	SendResponse(w, author, http.StatusOK)
}


// @Summary Update an author
// @Tags Authors
// @Param author_uuid path string true "Author UUID"
// @Param message body AuthorRequest true "Author"
// @Success 200 {object} repository.Author
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /authors/{author_uuid}/ [put]
func (h *CatalogHandler) authorUpdate(w http.ResponseWriter, r *http.Request) {
	author, err := h.getAuthorFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&author); err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	if err = h.repo.Author.Update(author); err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, author, http.StatusOK)
}

// @Summary Delete an author
// @Tags Authors
// @Success 200 {string} string "Empty response"
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Param author_uuid path string true "Author UUID"
// @Router /authors/{author_uuid}/ [delete]
func (h *CatalogHandler) authorDelete(w http.ResponseWriter, r *http.Request) {
	author, err := h.getAuthorFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	if err = h.repo.Author.Delete(author); err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}