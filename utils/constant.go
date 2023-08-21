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
	GET_IMAGE_BY_ID = "SELECT id, path, created_at, updated_at FROM images WHERE id = $1"
	INSERT_IMAGE    = "INSERT INTO images (id, path, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	DELETE_IMAGE    = "DELETE FROM images WHERE id = $1"
	UPDATE_IMAGE    = "UPDATE images SET path = $2, updated_at = $3 WHERE id = $1"

	//MEMBER
	GET_MEMBER_BY_ID       = "SELECT id, name, phone_no, no_identity, image_path, loan_status, created_at, updated_at, created_by FROM member WHERE id = $1"
	GET_ALL_MEMBER         = "SELECT id, name, phone_no, no_identity, image_path, loan_status, created_at, updated_at, created_by FROM member"
	GET_MEMBER_BY_PHONENO  = "SELECT id, name, phone_no, no_identity, image_path, loan_status, created_at, updated_at, created_by FROM member WHERE phone_no = $1"
	GET_MEMBER_BY_IDMEMBER = "SELECT id, name, phone_no, no_identity, image_path, loan_status, created_at, updated_at, created_by FROM member WHERE no_identity = $1"
	INSERT_MEMBER          = "INSERT INTO member (id, name, phone_no, no_identity, image_path, loan_status, created_at, updated_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	DELETE_MEMBER          = "DELETE FROM member WHERE id = $1"
	UPDATE_MEMBER          = "UPDATE images SET name = $2, phone_no = $3, no_identity = $4, image_path = $5, loan_status = $6, created_at = $7, updated_at = $8, created_by = $9 WHERE id = $1"

	//BOOKLOAN
	INSERT_BOOKLOAN        = "INSERT INTO book_loaning (id, member_id, book_id, start_date, end_date, loan_status) VALUES ($1, $2, $3, $4, $5, $6)"
	UPDATE_STOCK_BOOKLOAN  = "UPDATE book_list SET stock = stock - 1 WHERE id = ?"
	UPDATE_MEMBER_BOOKLOAN = "UPDATE member SET loan_status = ? WHERE id = ?"
	GET_BOOKLOAN_BY_ID     = `SELECT m.id, m.name, m.loan_status FROM book_loan AS bl
	JOIN member AS m ON bl.member_id = m.id
	WHERE bl.member_id = $1`
	GET_BOOKLOAN_BY_ID_DETAIL = `SELECT b.book_name, bl.start_date, bl.end_date, bl.loan_status FROM book_loan AS bl
	JOIN book_list AS b ON bl.book_id = b.id
	WHERE bl.member_id = $1`
	GET_ALL_BOOKLOAN = `SELECT m.id, m.name, m.loan_status, b.book_name, bl.start_date, bl.end_date, bl.loan_status 
	FROM book_loan AS bl
	JOIN member AS m ON bl.member_id = m.id
	JOIN book_list AS b ON bl.book_id = b.id`
	UPDATE_BOOKLOAN_STATUS = "UPDATE book_loan SET loan_status = $2 WHERE id = $1"
)