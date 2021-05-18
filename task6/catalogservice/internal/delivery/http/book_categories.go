package http

import (
	"fmt"
	"net/http"
)

// @Summary Create a book-category relation
// @Tags Books
// @Param book_uuid path string true "Book UUID"
// @Param category_uuid path string true "Category UUID"
// @Success 201 {string} string "Empty response"
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /books/{book_uuid}/categories/{category_uuid}/ [post]
func (h *CatalogHandler) bookCategoriesCreate(w http.ResponseWriter, r *http.Request) {
	book, err := h.getBookFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	category, err := h.getCategoryFromRequest(r)
	if err != nil {
		SendResponse(w, err,  http.StatusNotFound)
		fmt.Println("categ not found")
		return
	}

	err = h.repo.Book.AddCategory(book, category)
	if err != nil {
		SendResponse(w, err,  http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// @Summary Delete a book-category relation
// @Tags Books
// @Param book_uuid path string true "Book UUID"
// @Param category_uuid path string true "Category UUID"
// @Success 200 {string} string "Empty response"
// @Failure 404 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /books/{book_uuid}/categories/{category_uuid}/ [delete]
func (h *CatalogHandler) bookCategoriesDelete(w http.ResponseWriter, r *http.Request) {
	book, err := h.getBookFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	category, err := h.getCategoryFromRequest(r)
	if err != nil {
		SendResponse(w, err, http.StatusNotFound)
		return
	}

	err = h.repo.Book.RemoveCategory(book, category)
	if err != nil {
		SendResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

