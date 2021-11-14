package query

const (
	QueryCreateProduct = `
	INSERT INTO products(id,
		product_name,
		price,
		sku,
		picture,
		created_at,
		updated_at,
		outlet_id
	)
	VALUES(?,?,?,?,?,?,?,?)
	`

	QueryCreateImage = `
	INSERT INTO image_products(id,
		display_image,
		product_id
	)
	VALUES(?,?,?)
	`

	QueryFindAllProduct = `
	SELECT
      	id,
      	product_name,
		price,
		sku,
		picture,
		outlet_id
    FROM
    	products
    ORDER BY created_at DESC
	`

	QueryFindProductById = `
	SELECT 
		id,
      	product_name,
		price,
		sku,
		picture,
		outlet_id 
	FROM 
		products 
	WHERE 
		id = ?
	`

	QueryUpdateProductByID = `
	UPDATE 
		products 
	SET 
		product_name = ?,
		price = ?,
		sku = ?,
		picture = ?,
		outlet_id = ?,
		updated_at = ?
	WHERE 
		id = ?
	`

	QueryDeleteProductById = `
	DELETE FROM 
		products 
	WHERE 
		id = ?
	`
)
