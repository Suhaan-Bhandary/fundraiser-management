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
		organizer.name as organizer_name, fundraiser.image_url, fundraiser.video_url, 
		fundraiser.target_amount, fundraiser.status, fundraiser.created_at, fundraiser.updated_at
	from fundraiser
	join organizer
	on fundraiser.organizer_id = organizer.id
	where fundraiser.id = $1;
`

const closeFundraiserQuery = `
	update fundraiser
	set status = 'inactive', updated_at = CURRENT_TIMESTAMP
	where id = $1;
`

const getOrganizerIdAndStatusFromFundraiserQuery = `
	select organizer_id, status
	from fundraiser
	where id = $1;
`

const getFundraiserStatusQuery = `
	select status
	from fundraiser
	where id = $1;
`

const banFundraiserQuery = `
	update fundraiser
	set status = 'banned', updated_at = CURRENT_TIMESTAMP
	where id = $1;
`

const unbanFundraiserQuery = `
	update fundraiser
	set status = 'active', updated_at = CURRENT_TIMESTAMP
	where id = $1;
`

const listFundraisersQuery = `
	select 
		fundraiser.id, fundraiser.title, fundraiser.description, fundraiser.organizer_id, 
		organizer.name as organizer_name, fundraiser.image_url, fundraiser.video_url, 
		fundraiser.target_amount, fundraiser.status, fundraiser.created_at, fundraiser.updated_at
	from fundraiser
	join organizer
	on fundraiser.organizer_id = organizer.id
	where	
		(
			fundraiser.title ilike '%' || $1 || '%' or 
			fundraiser.description ilike '%' || $1 || '%' or
			organizer.name ilike '%' || $1 || '%' 
		) and
		(CASE when $2 in ('active', 'inactive', 'banned') THEN fundraiser.status = $2 ELSE true END)
	order by 
	    CASE WHEN $3 = 'donation_id' and $4 THEN donation.id END ASC,
	    CASE WHEN $3 = 'donation_id' and not $4 THEN donation.id END DESC,

	    CASE WHEN $3 = 'fundraiser_id' and $4 THEN fundraiser.fundraiser_id END ASC,
	    CASE WHEN $3 = 'fundraiser_id' and not $4 THEN fundraiser.fundraiser_id END DESC,

	    CASE WHEN $3 = 'title' and $4 THEN fundraiser.title END ASC,
	    CASE WHEN $3 = 'title' and not $4 THEN fundraiser.title END DESC,

	    CASE WHEN $3 = 'amount' and $4 THEN donation.amount END ASC,
	    CASE WHEN $3 = 'amount' and not $4 THEN donation.amount END DESC,

	    CASE WHEN $3 = 'is_anonymous' and $4 THEN donation.is_anonymous END ASC,
	    CASE WHEN $3 = 'is_anonymous' and not $4 THEN donation.is_anonymous END DESC,

	    CASE WHEN $3 = 'created_at' and $4 THEN donation.created_at END ASC,
	    CASE WHEN $3 = 'created_at' and not $4 THEN donation.created_at END DESC,

	    CASE WHEN $3 = '' THEN donation.created_at END DESC
	offset $5 
	limit $6;
`

const updateFundraiserQuery = `
	update fundraiser
	set title = $1, description = $2, image_url = $3, video_url = $4, target_amount = $5 
	where id = $6;
`
