package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
	"net/http"
)

func (h *CatalogHandler) initBookRoutes(r *mux.Router) {
	listRouter := r.PathPrefix("/books").Subrouter()
	listRouter.HandleFunc("/", h.bookList).Methods("GET")
	listRouter.HandleFunc("/", h.bookCreate).Methods("POST")

	detailRouter := listRouter.PathPrefix("/{book_uuid:[a-f0-9-]+}").Subrouter()
	detailRouter.HandleFunc("/", h.bookUpdate).Methods("PUT", "PATCH")
	detailRouter.HandleFunc("/", h.bookDetail).Methods("GET")
	detailRouter.HandleFunc("/", h.bookDelete).Methods("DELETE")

	m2mRouter := detailRouter.PathPrefix("/categories/{category_uuid:[a-f0-9-]+}").Subrouter()
	m2mRouter.HandleFunc("/", h.bookCategoriesCreate).Methods("POST")
	m2mRouter.HandleFunc("/", h.bookCategoriesDelete).Methods("DELETE")
}

type BookRequest struct {
	Name        string    `json:"name"        example:"Clean code"`
	Author_uuid uuid.UUID `json:"author_uuid" example:"01234567-89ab-cdef-0123-456789abcdef"`
}

// @Summary Create a Book
// @Tags Books
// @Param message body BookRequest true "Book"
// @Success 201 {object} repository.Book
// @Failure 500 {object} ResponseError
// @Router /books/ [post]
func (h *CatalogHandler) bookCreate(w http.ResponseWriter, r *http.Request) {
	book := &repository.Book{UUID: uuid.New()}

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	book, err = h.repo.Book.Create(book)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, book, http.StatusCreated)
}

// @Summary Search books
// @Tags Books
// @Param page          query uint64 false "page number"
// @Param page_size     query uint64 false "page size"
// @Param categories    query string false "comma-separated uuids"
// @Param authors       query string false "comma-separated uuids"
// @Param category_name query string false "search by author name"
// @Param author_name   query string false "search by category name"
// @Success 200 {array} PaginatedResults{results=[]repository.Book}
// @Failure 500 {object} ResponseError
// @Router /books/ [get]
func (h *CatalogHandler) bookList(w http.ResponseWriter, r *http.Request) {
	filter := repository.BookFilter{Page: 1, PageSize: h.apiPageSize}
	if err := updateFilterFromRequest(r, &filter); err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	books, hasNext, err := h.repo.Book.Search(&filter)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = h.repo.Book.FetchCategories(books)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	pageData := PaginatedResults{
		Page: filter.Page,
		PageSize : filter.PageSize,
		Results: books,
	}
	if filter.Page > 1 {
		pageData.Previous = addPageToUrl(*r.URL, filter.Page - 1)
	}
	if hasNext {
		pageData.Next = addPageToUrl(*r.URL, filter.Page + 1)
	}

	SendResponse(w, pageData, http.StatusOK)
}

// @Summary Get a books
// @Tags Books
// @Param book_uuid path string true "Book UUID"
// @Success 200 {array} repository.Book
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /books/{book_uuid}/ [get]
func (h *CatalogHandler) bookDetail(w http.ResponseWriter, r *http.Request) {
	book, err := h.getBookFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	err = h.repo.Book.FetchCategories([]*repository.Book{book})
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, book, http.StatusOK)
}

// @Summary Update a books
// @Tags Books
// @Param book_uuid path string true "Book UUID"
// @Param message body BookRequest true "Book"
// @Success 200 {array} repository.Book
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /books/{book_uuid}/ [put]
func (h *CatalogHandler) bookUpdate(w http.ResponseWriter, r *http.Request) {
	book, err := h.getBookFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = h.repo.Book.Update(book)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	SendResponse(w, book, http.StatusOK)
}

// @Summary Delete a books
// @Tags Books
// @Success 200 {array} repository.Book
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Param book_uuid path string true "Book UUID"
// @Router /books/{book_uuid}/ [delete]
func (h *CatalogHandler) bookDelete(w http.ResponseWriter, r *http.Request) {
	book, err := h.getBookFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	err = h.repo.Book.Delete(book)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
