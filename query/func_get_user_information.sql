CREATE OR REPLACE FUNCTION get_user_information(
	p_user_id INT
)
RETURNS TABLE (
	user_name TEXT
	, role_id INT
	, role_name VARCHAR(255)
	, is_active BOOLEAN
)
AS $$
	SELECT 
		users.name AS user_name
		, users.role_id
		, roles.name AS role_name
		, users.is_active
	FROM 
		users
		JOIN roles
			ON roles.id = users.role_id
	WHERE 
		users.id = p_user_id;
$$ LANGUAGE sql