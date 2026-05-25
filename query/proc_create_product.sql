/*
	CALL create_product(
		'Pan'
		, 5
		, 'testpath'
	);
*/

CREATE OR REPLACE PROCEDURE create_product(
	p_name TEXT
	, p_quantity INT
	, p_image_path VARCHAR(255)
)
AS $$
BEGIN
	INSERT INTO products(name, quantity, image_path)
	VALUES(p_name, p_quantity, p_image_path)
;END;
$$ LANGUAGE plpgsql