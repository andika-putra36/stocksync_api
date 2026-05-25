CREATE TABLE transaction_items(
	id SERIAL PRIMARY KEY
	, transaction_id INT NOT NULL REFERENCES transactions(id)
	, product_id INT NOT NULL REFERENCES products(id)
	, quantity INT NOT NULL
	, created_at TIMESTAMP DEFAULT NOW()
	, updated_at TIMESTAMP DEFAULT NOW()
);