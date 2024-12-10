package main

type DB struct {
	// TODO: ...
}

type SavedPost struct {
	ID string
	Post
}

func NewDB() (*DB, error) {
	// TODO: sql lite file setup
	return &DB{}, nil
}

func (db DB) InsertPost(sp SavedPost) error {
	// TODO: ...
	return nil
}

func (db DB) GetPostsByExternalID(externalID string) ([]SavedPost, error) {
	// TODO: ...
	return []SavedPost{}, nil
}
