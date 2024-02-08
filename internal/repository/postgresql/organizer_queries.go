package repository

const insertOrganizerQuery = `
	insert into organizer (organization, detail, email, password, mobile) 
	values($1, $2, $3, $4, $5);
`

const getOrganizerIdPasswordQuery = `
	select id, password
	from organizer
	where email = $1;
`

const verifyOrganizerQuery = `
	update organizer
	set is_verified = true
	where id = $1;
`
