package repository

const insertFundraiserQuery = `
	insert into fundraiser (title, description, organizer_id, image_url, video_url, target_amount, status) 
	values($1, $2, $3, $4, $5, $6, $7) Returning id;
`

const deleteFundraiserQuery = `
	delete
	from fundraiser
	where id = $1;
`

const getOrganizerIdFromFundraiser = `
	select organizer_id
	from fundraiser
	where id = $1;
`