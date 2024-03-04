package repository

const insertOrganizerQuery = `
	insert into organizer (name, detail, email, password, mobile) 
	values($1, $2, $3, $4, $5);
`

const getOrganizerIdPasswordQuery = `
	select id, password, is_verified
	from organizer
	where email = $1;
`

const deleteOrganizerQuery = `
	delete
	from organizer
	where id = $1;
`

const verifyOrganizerQuery = `
	update organizer
	set is_verified = true, updated_at = CURRENT_TIMESTAMP
	where id = $1;
`

const listOrganizersQuery = `
	select id, name, detail, email, mobile, is_verified
	from organizer
	where 
		(name ilike '%' || $1 || '%' or detail ilike '%' || $1 || '%') and
		(CASE when $2 in ('true', 'false') THEN is_verified = ($2 = 'true') ELSE true END)
	order by 
	    CASE WHEN $3 = 'id' and $4 THEN id END ASC,
	    CASE WHEN $3 = 'id' and not $4 THEN id END DESC,

	    CASE WHEN $3 = 'name' and $4 THEN name END ASC,
	    CASE WHEN $3 = 'name' and not $4 THEN name END DESC,

	    CASE WHEN $3 = 'email' and $4 THEN email END ASC,
	    CASE WHEN $3 = 'email' and not $4 THEN email END DESC,

	    CASE WHEN $3 = 'mobile' and $4 THEN mobile END ASC,
	    CASE WHEN $3 = 'mobile' and not $4 THEN mobile END DESC,

	    CASE WHEN $3 = 'is_verified' and $4 THEN is_verified END ASC,
	    CASE WHEN $3 = 'is_verified' and not $4 THEN is_verified END DESC,

	    CASE WHEN $3 = 'created_at' and $4 THEN created_at END ASC,
	    CASE WHEN $3 = 'created_at' and not $4 THEN created_at END DESC,

	    CASE WHEN $3 = 'updated_at' and $4 THEN updated_at END ASC,
	    CASE WHEN $3 = 'updated_at' and not $4 THEN updated_at END DESC,

	    CASE WHEN $3 = '' THEN created_at END DESC
	offset $5 
	limit $6;
`

const getOrganizersListCountQuery = `
	select count(*) 
	from organizer
	where	(name ilike '%' || $1 || '%' or detail ilike '%' || $1 || '%') and
		(CASE when $2 in ('true', 'false') THEN is_verified = ($2 = 'true') ELSE true END);
`

const getOrganizersWithFilter = `
	select id, name, detail, email, mobile, is_verified
	from organizer
	where (name ilike '%' || $1 || '%') 
	order by created_at desc;
`

const getOrganizerQuery = `
	select id, name, detail, email, mobile, is_verified
	from organizer
	where id = $1;
`

const updateOrganizerQuery = `
	update organizer
	set email = $1, detail = $2, mobile = $3, updated_at = CURRENT_TIMESTAMP 
	where id = $4;
`
