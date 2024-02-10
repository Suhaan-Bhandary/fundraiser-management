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

const getOrganizers = `
	select id, name, detail, email, mobile, is_verified
	from organizer
	order by created_at desc;
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
