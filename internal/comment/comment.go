package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchingComment = errors.New("Failed to fetch the id for the comment")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// All the methods, required for service to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

type Service struct {
	Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retreving a comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorFetchingComment
	}

	return cmt, nil
}
