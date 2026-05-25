/*
	SELECT * FROM get_user_products(1);
*/

CREATE OR REPLACE FUNCTION get_user_products(
	p_user_id INT
)
RETURNS TABLE (
	id INT
	, name TEXT
	, quantity INT
	, image_path VARCHAR(255)
)
AS $$
	SELECT
		products.id
		, products.name
		, SUM(transaction_items.quantity) AS quantity
		, products.image_path
	FROM 
		transaction_items
		JOIN transactions 
			ON transactions.id = transaction_items.transaction_id
		JOIN products 
			ON products.id = transaction_items.product_id
	WHERE transactions.user_id = p_user_id
	GROUP BY 
		products.id
		, products.name
		, products.image_path
	ORDER BY products.id
;
$$ LANGUAGE sql