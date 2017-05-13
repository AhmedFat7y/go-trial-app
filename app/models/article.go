package models

import (
	"database/sql"
	"log"
)

// Article struct is a model for the flattened articles table
type Article struct {
	ID           int64         `db:"id, primarykey, autoincrement"`
	Name         string        `db:"name"`
	Content      string        `db:"content"`
	CategoryID   int64         `db:"category_id"`
	CategoryName string        `db:"category_name"`
	CategorySlug string        `db:"category_slug"`
	AuthorID     int64         `db:"author_id"`
	AuthorName   string        `db:"author_name"`
	AuthorAbout  string        `db:"author_about"`
	CreatedAt    sql.NullInt64 `db:"created_at"`
	UpdatedAt    sql.NullInt64 `db:"updated_at"`
}

// GetArticles gets employees within offset, and limit
func GetArticles(offset int, limit int) (bool, []Article) {
	var articles []Article
	_, err := dbmap.Select(&articles, "Select id, name, content, category_id, category_name, category_slug, author_id, author_name, author_about, created_at, updated_at FROM flat_articles LIMIT ? OFFSET ?", limit, offset)

	if err != nil {
		log.Print("ERROR get articles: ")
		log.Println(err)
	}

	return (err == nil), articles
}
