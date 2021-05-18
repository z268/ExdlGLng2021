package repository

type CatalogRepository struct {
	Book     BookRepository
	Author   AuthorRepository
	Category CategoryRepository
}
