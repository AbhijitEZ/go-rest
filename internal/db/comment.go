package db

import (
	"context"
	"database/sql"
	"fmt"
	"go-rest/internal/comment"

	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

// coverting the SQL data-types to actual Comment Struct
func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Author.String,
	}
}

func (db *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	row := db.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author
			FROM comments WHERE id = $1`, uuid,
	)
	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to fetch comment %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (db *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	row, err := db.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
		(id, slug, author, body) VALUES
		(:id, :slug, :author, :body)`, postRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert the error %w", err)
	}
	if err := row.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close the comment %w", err)
	}

	return cmt, nil
}

func (db *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := db.Client.ExecContext(
		ctx,
		`DELETE FROM comments WHERE id = $1`, id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete the comment %w", err)
	}
	return nil
}

func (db *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	row, err := db.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		slug = :slug,
		author = :author
		body = :body
		WHERE id = :id`,
		cmtRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to update the comment %w", err)
	}

	if err := row.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close the rows %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}
