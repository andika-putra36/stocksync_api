/*
	CALL register_user(
		1
		, 'Elvin'
		, 'elvin@mail.test'
		, 'hashpassword3'
	);
*/

CREATE OR REPLACE PROCEDURE register_user(
	p_role_id INT
	, p_name VARCHAR(255)
	, p_email VARCHAR(255)
	, p_password_hash TEXT
)
AS $$
BEGIN
	INSERT INTO users(role_id, name, email, password_hash)
	VALUES(p_role_id, p_name, p_email, p_password_hash)
;END;
$$ LANGUAGE plpgsql