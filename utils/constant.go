package utils

const (
	//USER
	INSERT_USER             = "INSERT INTO user_credential(id, username, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5)"
	UPDATE_USER             = "UPDATE user_credential SET username = $1, password = $2, updated_at = $3 WHERE id = $4"
	SELECT_USER_BY_USERNAME = "SELECT id, username, password, created_at, updated_at FROM user_credential WHERE username = $1"
	DELETE_USER             = "DELETE FROM user_credential WHERE id = $1"
	SELECT_USER_BY_ID       = "SELECT id, username, password, created_at, updated_at FROM user_credential WHERE id = $1"

	//BOOK
	GET_ALL_BOOK      = "SELECT id, book_name, created_at, created_by, stock FROM book_list"
	GET_BOOK_BY_ID    = "SELECT id, book_name, created_at, created_by, stock FROM book_list = $1"
	INSERT_BOOK       = "INSERT INTO book_list (id, book_name, created_at, created_by, stock) VALUES ($1, $2, $3, $4, $5)"
	DELETE_BOOK       = "DELETE FROM book_list WHERE id=$1"
	UPDATE_BOOK_STOCK = "UPDATE book_list SET stock = $2 WHERE id = $1"

	//IMAGE
	GET_IMAGE_BY_ID = "SELECT id, path, created_at, updated_at FROM images = $1"
	INSERT_IMAGE    = "INSERT INTO images (id, path, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	DELETE_IMAGE    = "DELETE FROM images WHERE id = $1"
	UPDATE_IMAGE    = "UPDATE images SET path = $2, updated_at = $3 WHERE id = $1"
)