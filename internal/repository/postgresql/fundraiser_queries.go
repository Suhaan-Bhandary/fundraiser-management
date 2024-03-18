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
		organizer.name as organizer_name, fundraiser.image_url, fundraiser.video_url, fundraiser.target_amount,
        (
            select SUM(donation.amount)
            from donation
            where donation.fundraiser_id = $1
        ) as amount_collected,
        fundraiser.status, fundraiser.created_at, fundraiser.updated_at
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
		organizer.name as organizer_name, fundraiser.image_url, fundraiser.video_url, fundraiser.target_amount,
        (
            select SUM(donation.amount)
            from donation
            where donation.fundraiser_id = fundraiser.id
        ) as amount_collected,
        fundraiser.status, fundraiser.created_at, fundraiser.updated_at
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
	    CASE WHEN $3 = 'fundraiser_id' and $4 THEN fundraiser.id END ASC,
	    CASE WHEN $3 = 'fundraiser_id' and not $4 THEN fundraiser.id END DESC,

	    CASE WHEN $3 = 'organizer_id' and $4 THEN fundraiser.organizer_id END ASC,
	    CASE WHEN $3 = 'organizer_id' and not $4 THEN fundraiser.organizer_id END DESC,

	    CASE WHEN $3 = 'title' and $4 THEN fundraiser.title END ASC,
	    CASE WHEN $3 = 'title' and not $4 THEN fundraiser.title END DESC,

	    CASE WHEN $3 = 'description' and $4 THEN fundraiser.description END ASC,
	    CASE WHEN $3 = 'description' and not $4 THEN fundraiser.description END DESC,

	    CASE WHEN $3 = 'organizer_name' and $4 THEN organizer.name END ASC,
	    CASE WHEN $3 = 'organizer_name' and not $4 THEN organizer.name  END DESC,

	    CASE WHEN $3 = 'target_amount' and $4 THEN fundraiser.target_amount END ASC,
	    CASE WHEN $3 = 'target_amount' and not $4 THEN fundraiser.target_amount END DESC,

	    CASE WHEN $3 = 'status' and $4 THEN fundraiser.status END ASC,
	    CASE WHEN $3 = 'status' and not $4 THEN fundraiser.status END DESC,

	    CASE WHEN $3 = 'created_at' and $4 THEN fundraiser.created_at END ASC,
	    CASE WHEN $3 = 'created_at' and not $4 THEN fundraiser.created_at END DESC,

	    CASE WHEN $3 = 'updated_at' and $4 THEN fundraiser.updated_at END ASC,
	    CASE WHEN $3 = 'updated_at' and not $4 THEN fundraiser.updated_at END DESC,

	    CASE WHEN $3 = '' THEN fundraiser.created_at END DESC
	offset $5 
	limit $6;
`

const getListFundraisersCountQuery = `
	select count(fundraiser.id)
	from fundraiser
	join organizer
	on fundraiser.organizer_id = organizer.id
	where	
		(
			fundraiser.title ilike '%' || $1 || '%' or 
			fundraiser.description ilike '%' || $1 || '%' or
			organizer.name ilike '%' || $1 || '%' 
		) and
		(CASE when $2 in ('active', 'inactive', 'banned') THEN fundraiser.status = $2 ELSE true END);
`

const updateFundraiserQuery = `
	update fundraiser
	set title = $1, description = $2, image_url = $3, video_url = $4, target_amount = $5 
	where id = $6;
`
