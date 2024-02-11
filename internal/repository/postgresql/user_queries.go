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

const deleteUserQuery = `
	delete
	from users
	where id = $1;
`

const listUsersQuery = `
	select id, first_name, last_name, email
	from users
	where (first_name ilike '%' || $1 || '%' or last_name ilike '%' || $1 || '%' or email ilike '%' || $1 || '%')
	order by 
	    CASE WHEN $2 = 'id' and $3 THEN id END ASC,
	    CASE WHEN $2 = 'id' and not $3 THEN id END DESC,

	    CASE WHEN $2 = 'first_name' and $3 THEN first_name END ASC,
	    CASE WHEN $2 = 'first_name' and not $3 THEN first_name END DESC,

	    CASE WHEN $2 = 'last_name' and $3 THEN last_name END ASC,
	    CASE WHEN $2 = 'last_name' and not $3 THEN last_name END DESC,

	    CASE WHEN $2 = 'email' and $3 THEN email END ASC,
	    CASE WHEN $2 = 'email' and not $3 THEN email END DESC,

	    CASE WHEN $2 = 'created_at' and $3 THEN created_at END ASC,
	    CASE WHEN $2 = 'created_at' and not $3 THEN created_at END DESC,

	    CASE WHEN $2 = 'updated_at' and $3 THEN updated_at END ASC,
	    CASE WHEN $2 = 'updated_at' and not $3 THEN updated_at END DESC,

	    CASE WHEN $2 = '' THEN created_at END DESC
	offset $4 
	limit $5;
`

const getListUsersCountQuery = `
	select count(id) as count
	from users
	where (first_name ilike '%' || $1 || '%' or last_name ilike '%' || $1 || '%' or email ilike '%' || $1 || '%');
`

const getUserQuery = `
	select id, first_name, last_name, email
	from users
	where id = $1;
`
