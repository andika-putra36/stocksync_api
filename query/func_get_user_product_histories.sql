/*
	SELECT * FROM get_user_product_histories(1, 1);
*/

CREATE OR REPLACE FUNCTION get_user_product_histories(
	p_user_id INT
	, p_product_id INT
)
RETURNS TABLE (
	transaction_id INT
	, quantity INT
	, created_at TIMESTAMP
	, updated_at TIMESTAMP
)
AS $$
	SELECT
		transaction_items.transaction_id
		, transaction_items.quantity
		, transaction_items.created_at
		, transaction_items.updated_at
	FROM 
		transaction_items
		JOIN transactions 
			ON transactions.id = transaction_items.transaction_id
		JOIN products 
			ON products.id = transaction_items.product_id
	WHERE 
		transactions.user_id = p_user_id
		AND transaction_items.product_id = p_product_id
;
$$ LANGUAGE sql