CREATE TABLE user_credential (
	id VARCHAR(100) PRIMARY KEY NOT NULL,
	username VARCHAR(50) NOT NULL,
	password VARCHAR NOT NULL,
	created_at TIMESTAMP, 
	updated_at TIMESTAMP 
);

CREATE TABLE book_list (
	id VARCHAR(100) PRIMARY KEY NOT NULL,
	book_name VARCHAR(50) NOT NULL,
	stock INTEGER NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE member (
	id VARCHAR(100) PRIMARY KEY NOT NULL,
	name VARCHAR(50) NOT NULL,
	phone_no VARCHAR(50) NOT NULL,
	no_identity VARCHAR(50) NOT NULL,
	photo BYTEA NOT NULL,
	loan_status VARCHAR(50) NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	created_by VARCHAR(50) NOT NULL
);

CREATE TABLE book_loaning (
	id VARCHAR(100) PRIMARY KEY NOT NULL,
	member_id VARCHAR(100) NOT NULL,
	book_id VARCHAR NOT NULL,
	start_date DATE NOT NULL,
	end_date DATE NOT NULL,
	late_charge_day INTEGER NOT NULL,
	late_charge FLOAT NOT NULL,
	loan_status VARCHAR(50)
);

