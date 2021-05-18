package http

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (h *CatalogHandler) getBookFromRequest(r *http.Request) (*repository.Book, error) {
	id, err := uuid.Parse(mux.Vars(r)["book_uuid"])
	if err != nil {
		return nil, err
	}

	return h.repo.Book.Get(id)
}

func (h *CatalogHandler) getAuthorFromRequest(r *http.Request) (*repository.Author, error) {
	id, err := uuid.Parse(mux.Vars(r)["author_uuid"])
	if err != nil {
		return nil, err
	}

	return h.repo.Author.Get(id)
}

func (h *CatalogHandler) getCategoryFromRequest(r *http.Request) (*repository.Category, error) {
	id, err := uuid.Parse(mux.Vars(r)["category_uuid"])
	if err != nil {
		return nil, err
	}

	return h.repo.Category.Get(id)
}

// Parse GET arguments to BookFilter
func updateFilterFromRequest(r *http.Request, filter *repository.BookFilter) error {
	urlParams := r.URL.Query()

	pageStr := urlParams.Get("page")
	if page, err := strconv.Atoi(pageStr); err == nil {
		filter.Page = uint64(page)
	} else if len(pageStr) > 0 {
		return err
	}

	perPageStr := urlParams.Get("page_size")
	if page, err := strconv.Atoi(perPageStr); err == nil {
		filter.PageSize = uint64(page)
	} else if len(perPageStr) > 0 {
		return err
	}

	for _, catStr := range(strings.Split(urlParams.Get("categories"), ",")) {
		if category_uuid, err := uuid.Parse(catStr); err == nil {
			filter.Categories = append(filter.Categories, category_uuid)
		} else if len(catStr) > 0 {
			return err
		}
	}

	for _, authorStr := range(strings.Split(urlParams.Get("authors"), ",")) {
		if author_uuid, err := uuid.Parse(authorStr); err == nil {
			filter.Authors = append(filter.Authors, author_uuid)
		} else if len(authorStr) > 0 {
			return err
		}
	}

	filter.AuthorName = urlParams.Get("author_name")
	filter.CategoryName = urlParams.Get("category_name")

	return nil
}

func addPageToUrl(url url.URL, page uint64) string {
	query := url.Query()
	query.Set("page", fmt.Sprint(page))
	url.RawQuery = query.Encode()
	return url.String()
}