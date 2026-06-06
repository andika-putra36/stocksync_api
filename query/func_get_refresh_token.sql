CREATE OR REPLACE FUNCTION get_refresh_token(
	p_token TEXT
)
RETURNS TABLE(
	id INT
	, user_id INT
	, email VARCHAR(255)
	, token TEXT
	, expired_at TIMESTAMP
)
AS $$

	SELECT 
		refresh_tokens.id
		, refresh_tokens.user_id
		, users.email
		, refresh_tokens.token
		, refresh_tokens.expired_at
	FROM
		refresh_tokens
		JOIN users
			ON users.id = refresh_tokens.user_id
	WHERE 
		refresh_tokens.token = p_token;
$$ LANGUAGE sql;