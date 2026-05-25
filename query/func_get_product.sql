/*
	SELECT * FROM get_product(2);
*/

CREATE OR REPLACE FUNCTION get_product(
	p_id INT
)
RETURNS TABLE (
	id INT
	, name TEXT
	, quantity INT
	, image_path VARCHAR(255)
	, is_delete BOOLEAN
	, created_at TIMESTAMP
	, updated_at TIMESTAMP
)
AS $$
	SELECT
		products.id
		, products.name
		, products.quantity
		, products.image_path
		, products.is_delete
		, products.created_at
		, products.updated_at
	FROM products
	WHERE 
		products.id = p_id
		AND products.is_delete = false
;
$$ LANGUAGE sql