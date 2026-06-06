CREATE OR REPLACE PROCEDURE delete_refresh_token(
	p_user_id 	INT
)
LANGUAGE plpgsql
AS $$
BEGIN
	DELETE FROM refresh_tokens
	WHERE user_id = p_user_id;
END;
$$;