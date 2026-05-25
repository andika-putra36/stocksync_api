CREATE TABLE refresh_tokens(
	id SERIAL PRIMARY KEY
	, user_id INT REFERENCES users(id)
	, token TEXT NOT NULL UNIQUE
	, expires_at TIMESTAMP NOT NULL
	, created_at TIMESTAMP DEFAULT NOW()
	, updated_at TIMESTAMP DEFAULT NOW()
);