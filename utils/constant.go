package utils

const (
	//USER
	INSERT_USER             = "INSERT INTO user_credential(id, username, password, created_at, updated_at) VALUES($1, $2, $3, $4)"
	UPDATE_USER             = "UPDATE user_credential SET username = $1, password = $2, updated_at = $3 WHERE id = $5"
	SELECT_USER_BY_USERNAME = "SELECT id, username, password, created_at, updated_at FROM user_credential WHERE username = $1"
	DELETE_USER             = "DELETE FROM user_credential WHERE id = $1"
)