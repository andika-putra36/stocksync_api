/*
	CALL update_product(
		1
		, 'Frying Pan'
		, 10
		, 'testpathedited'
	)
*/

CREATE OR REPLACE PROCEDURE update_product(
	p_id INT
	, p_name TEXT
	, p_quantity INT
	, p_image_path VARCHAR(255)
)
AS $$
BEGIN
	UPDATE products
	SET 
		name = p_name
		, quantity = p_quantity
		, image_path = p_image_path
		, updated_at = NOW()
	WHERE id = p_id
;END;
$$ LANGUAGE plpgsql