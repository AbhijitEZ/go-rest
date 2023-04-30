package db

import (
	"context"
	"database/sql"
	"fmt"
	"go-rest/internal/comment"
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
