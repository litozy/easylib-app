CREATE TABLE book_loaning (
	id VARCHAR(100) PRIMARY KEY NOT NULL,
	member_id VARCHAR(100) NOT NULL,
	book_id VARCHAR(100) NOT NULL,
	start_date DATE NOT NULL,
	end_date DATE NOT NULL,
	late_charge FLOAT NOT NULL,
	loan_status VARCHAR(50),
	FOREIGN KEY (member_id) REFERENCES member(id),
	FOREIGN KEY (book_id) REFERENCES book_id(id)
);