CREATE TABLE user_credential (
	id VARCHAR(100) PRIMARY KEY NOT NULL,
	username VARCHAR(50) NOT NULL,
	name VARCHAR(50) NOT NULL,
	password VARCHAR NOT NULL,
	created_at TIMESTAMP, 
	updated_at TIMESTAMP 
);