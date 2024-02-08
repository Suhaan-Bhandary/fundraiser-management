package repository

const insertUserQuery = `
	insert into users (first_name, last_name, email, password) 
	values($1, $2, $3, $4);
`

const getUserPasswordQuery = `
	select id, password
	from users
	where email = $1;
`

const getUsersQuery = `
	select id, first_name, last_name, email
	from users;
`
