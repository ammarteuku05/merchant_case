package query

const (
	FindOutletUserByID = `
SELECT * 
FROM 
	outlets 
WHERE 
	id = ?
`

	CreateOutletbyUser = `
INSERT INTO outlets(id,
	outlet_name,
	picture,
	user_id,
	created_at,
	updated_at
)
VALUES(?,?,?,?,?,?)
`

	GetAllOutlets = `
SELECT
	  id,
	  outlet_name,
	  picture,
	user_id
FROM
	outlets
ORDER BY created_at DESC
`
)
