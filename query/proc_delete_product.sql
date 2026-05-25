/*
	CALL delete_product(1)
*/

CREATE OR REPLACE PROCEDURE delete_product(
	p_id INT
)
AS $$
BEGIN
	UPDATE products
	SET
		is_delete = TRUE
		, updated_at = NOW()
	WHERE id = p_id
;END;
$$ LANGUAGE plpgsql