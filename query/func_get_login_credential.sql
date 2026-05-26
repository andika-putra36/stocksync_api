/*
	SELECT * FROM get_login_credential('andika@mail.test');
*/

CREATE OR REPLACE FUNCTION get_login_credential(
	p_email VARCHAR(255)
)
RETURNS TABLE (
	role_id INT
	, email VARCHAR(255)
	, password_hash TEXT
)
AS $$
	SELECT 
		role_id
		, email
		, password_hash
	FROM users
	WHERE 
		email = p_email
;
$$ LANGUAGE sql