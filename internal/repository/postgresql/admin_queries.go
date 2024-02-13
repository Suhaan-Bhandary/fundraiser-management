package repository

const insertAdminQuery = `
	INSERT INTO admin 
	(username, password) 
	VALUES
	($1, $2)
`

const getAdminIdPasswordQuery = `
	select id, password
	from admin
	where username = $1;
`
