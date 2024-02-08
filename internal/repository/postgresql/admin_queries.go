package repository

const getAdminIdPasswordQuery = `
	select id, password
	from admin
	where username = $1;
`
