package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	PostsTableName string = "POSTS"
	DriverName     string = "sqlite3"
	DataSourceName string = "./data.db"
)

type DB struct {
	Client         *sql.DB
	driverName     string
	dataSourceName string
}

type SavedPost struct {
	ID        string // The ID of the post for use in this application
	TwitterID string // The ID of the twitter account this post is associated with
	Post
}

func NewDB(driverName string, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		create table if not exists %s (
			id text not null primary key,
			twitter_id text not null,
			url text not null unique,
			title text not null
		);
	`, PostsTableName)

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client:         db,
		driverName:     driverName,
		dataSourceName: dataSourceName,
	}, nil
}

func (db DB) InsertSavedPost(sp SavedPost) error {
	row := db.Client.QueryRow(fmt.Sprintf(`select * from %s where url = "%s"`, PostsTableName, sp.Url))
	if row.Scan() != sql.ErrNoRows {
		return fmt.Errorf(`SavedPost with url "%s" already present in db`, sp.Url)
	}

	insertQuery := fmt.Sprintf(`
		insert into %s (id, twitter_id, url, title) 
		values (?, ?, ?, ?);
	`, PostsTableName)

	if _, err := db.Client.Exec(insertQuery, sp.ID, sp.TwitterID, sp.Url, sp.Title); err != nil {
		return err
	}

	return nil
}

func (db DB) GetSavedPostsByTwitterID(twitterID string) ([]SavedPost, error) {
	var savedPosts []SavedPost

	rows, err := db.Client.Query(fmt.Sprintf(`select * from %s where twitter_id = "%s"`, PostsTableName, twitterID))
	if err != nil {
		return savedPosts, err
	}
	defer rows.Close()

	for rows.Next() {
		sp, err := scanIntoSavedPost(rows)
		if err != nil {
			return []SavedPost{}, err
		}
		savedPosts = append(savedPosts, *sp)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return []SavedPost{}, err
	}

	return savedPosts, nil
}

func scanIntoSavedPost(rows *sql.Rows) (*SavedPost, error) {
	sp := new(SavedPost)
	err := rows.Scan(
		&sp.ID,
		&sp.TwitterID,
		&sp.Url,
		&sp.Title,
	)
	return sp, err
}
