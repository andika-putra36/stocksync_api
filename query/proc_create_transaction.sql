/*
	CALL create_transaction(
		3
		, ARRAY[1, 2]
		, ARRAY[10, 4]
	);
*/

CREATE OR REPLACE PROCEDURE create_transaction(
	p_user_id INT
	, p_product_ids INT[]
	, p_quantities INT[]
)
AS $$
/* PRE PROCESS */
DECLARE 
	v_transaction_id INT;
	v_product_id INT;
	v_quantity INT;
	v_stock INT;
/* END OF PRE PROCESS */
BEGIN
	/* 
		VALIDATION PROCESS: 
		- Check if []product_ids and []quantities is the same
		- Check if every product is found
		- Check if every product is not deleted (skipped)
		- Check if every product id is available
	*/

	IF array_length(p_product_ids, 1) <> array_length(p_quantities, 1) THEN
	    RAISE EXCEPTION 'product_ids and quantities arrays must have the same length';
	END IF;

	PERFORM 1 FROM products
	WHERE products.id = ANY(p_product_ids)
	ORDER BY products.id
	FOR UPDATE;

	FOR v_product_id, v_quantity IN
    	SELECT pid, qty
		FROM UNNEST(p_product_ids, p_quantities) AS t(pid, qty)
		ORDER BY pid
	LOOP
		SELECT products.quantity INTO v_stock
		FROM products
		WHERE products.id = v_product_id;

		IF NOT FOUND THEN
		    RAISE EXCEPTION 'Product not found for product_id: %', v_product_id;
		END IF;

		IF v_stock < v_quantity THEN
			RAISE EXCEPTION 
				'Insufficient stock for product_id: %. Available: %, Requested: %'
				, v_product_id
				, v_stock
				, v_quantity;
		END IF;
	END LOOP;
	/* END OF VALIDATION PROCESS */
	
	/* COMMIT PROCESS */
	INSERT INTO transactions (user_id)
	VALUES (p_user_id)
	RETURNING id INTO v_transaction_id;

	FOR v_product_id, v_quantity IN
    	SELECT pid, qty
		FROM UNNEST(p_product_ids, p_quantities) AS t(pid, qty)
		ORDER BY pid
	LOOP
		UPDATE products
		SET quantity = products.quantity - v_quantity
		WHERE products.id = v_product_id;
		
		INSERT INTO transaction_items (transaction_id, product_id, quantity)
		VALUES (v_transaction_id, v_product_id, v_quantity);
	END LOOP;
	/* END OF COMMIT PROCESS */
	
EXCEPTION
	/* ROLLBACK PROCESS */
	WHEN OTHERS THEN
		RAISE EXCEPTION 'Transaction failed: %', SQLERRM;
	/* ROLLBACK PROCESS */
END;
$$ LANGUAGE plpgsql