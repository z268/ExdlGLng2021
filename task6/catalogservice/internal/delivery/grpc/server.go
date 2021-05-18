package grpc

import (
	"context"
	_ "github.com/go-sql-driver/mysql"

	"github.com/google/uuid"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewBookServiceServer(repo *repository.CatalogRepository) *bookServiceServer {
	return &bookServiceServer{repo: repo}
}

type bookServiceServer struct {
	UnimplementedBookServiceServer
	repo *repository.CatalogRepository
}

func (s *bookServiceServer) GetBookByUUID(ctx context.Context, in *GetBooksRequest) (*GetBooksResponse, error) {
	book_uuid, err := uuid.Parse(in.GetBookUuid()[0])
	if err != nil {
		return nil, status.Error(codes.NotFound, "id was not found")
	}
	book, err := s.repo.Book.Get(book_uuid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "book uuid was not found")
	}
	author, err := s.repo.Author.Get(book.Author_uuid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "author uuid was not found")
	}

	books := []*Book{
		{
			Name:   book.Name,
			Author: author.Name,
		},
	}
	return &GetBooksResponse{Book: books}, nil
}