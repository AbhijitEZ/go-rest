package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchingComment = errors.New("failed to fetch the id for the comment")
	ErrorPostComment     = errors.New("failed to create new comment")
	ErrorUpdateComment   = errors.New("failed to update comment")
	ErrorDeleteComment   = errors.New("failed to delete new comment")
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
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
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

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	cmtResult, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorPostComment
	}

	return cmtResult, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	err := s.Store.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return ErrorDeleteComment
	}
	return nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	result, err := s.Store.UpdateComment(ctx, id, cmt)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorUpdateComment
	}
	return result, nil
}
