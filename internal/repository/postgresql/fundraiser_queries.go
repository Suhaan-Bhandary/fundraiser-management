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

const getFundraiserQuery = `
	select 
		fundraiser.id, fundraiser.title, fundraiser.description, fundraiser.organizer_id, 
		organizer.organization as organizer_name, fundraiser.image_url, fundraiser.video_url, 
		fundraiser.target_amount, fundraiser.status, fundraiser.created_at, fundraiser.updated_at
	from fundraiser
	join organizer
	on fundraiser.organizer_id = organizer.id
	where fundraiser.id = $1;
`
