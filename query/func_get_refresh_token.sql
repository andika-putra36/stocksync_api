CREATE OR REPLACE FUNCTION get_refresh_token(
	p_token TEXT
)
RETURNS TABLE(
	id INT
	, user_id INT
	, token TEXT
	, expired_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	SELECT 
		refresh_tokens.id
		, refresh_tokens.user_id
		, refresh_tokens.token
		, refresh_tokens.expired_at
	FROM
		refresh_tokens
	WHERE 
		refresh_tokens.token = p_token;
END;
$$;