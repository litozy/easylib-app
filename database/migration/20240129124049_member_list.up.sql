CREATE TABLE member (
	id VARCHAR(100) PRIMARY KEY NOT NULL,
	name VARCHAR(50) NOT NULL,
	phone_no VARCHAR(50) NOT NULL,
	no_identity VARCHAR(50) NOT NULL,
	image_path TEXT NOT NULL,
	loan_status BOOLEAN,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	created_by VARCHAR(50) NOT NULL
);