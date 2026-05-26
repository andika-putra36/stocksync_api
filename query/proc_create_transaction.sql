/*
	CALL 
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
		- Check if every product id is available
	*/
	FOR v_product_id, v_quantity IN
    	SELECT pid, qty
		FROM UNNEST(p_product_ids, p_quantities) AS t(pid, qty)
		ORDER BY pid
	LOOP
		SELECT quantity INTO v_stock
		FROM products
		WHERE product_id = v_product_id
		FOR UPDATE;

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
	RETURNING transaction_id INTO v_transaction_id;

	FOR v_product_id, v_quantity IN
    	SELECT pid, qty
		FROM UNNEST(p_product_ids, p_quantities) AS t(pid, qty)
		ORDER BY pid
	LOOP
		UPDATE products
		SET quantity = quantity - v_quantity
		WHERE product_id = v_product_id;
		
		INSERT INTO transaction_items (transaction_id, product_id, quantity)
		VALUES (v_transaction_id, v_product_id, v_quantity);
	END LOOP;

	COMMIT;
	/* END OF COMMIT PROCESS */
	
EXCEPTION
	/* ROLLBACK PROCESS */
	WHEN OTHERS THEN
		ROLLBACK;
		RAISE EXCEPTION 'Transaction failed: %', SQLERRM;
	/* ROLLBACK PROCESS */
END;
$$ LANGUAGE plpgsql