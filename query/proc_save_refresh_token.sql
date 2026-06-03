CREATE OR REPLACE PROCEDURE save_refresh_token(
	p_user_id 	INT
	, p_token	TEXT
	, p_expired_at TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
	INSERT INTO refresh_tokens(user_id, token, expired_at)
	VALUES(p_user_id, p_token, p_expired_at);
END;
$$;