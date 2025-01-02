package database

import (
	"context"
	"database/sql"
	"log"
	"strings"
)

func ClearParagraphSymbolFromString(str string) (result string) {
	if isTrue := strings.Contains(str, "</p>"); isTrue {
		firstFilter := strings.ReplaceAll(str, "</p>", "")
		result = strings.ReplaceAll(firstFilter, "<p>", "")
		return
	}
	return str
}

func (q *Queries) InsertPostsBulk(ctx context.Context, db *sql.DB, posts []CreatePostParams) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("something was wrong initializing database transaction: %v", err)
		return nil
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	txQueries := q.WithTx(tx)
	log.Printf("Loading %d posts into table", len(posts))
	for _, value := range posts {
		newDescription := ClearParagraphSymbolFromString(value.Description)
		value.Description = newDescription
		err := txQueries.CreatePost(ctx, value)
		if err != nil {
			return err
		}
	}
	return nil
}
