CREATE TABLE book_list (
	id VARCHAR(100) PRIMARY KEY NOT NULL,
	book_name VARCHAR(50) NOT NULL,
	stock INTEGER NOT NULL,
	created_at TIMESTAMP,
	created_by VARCHAR(50) NOT NULL
);