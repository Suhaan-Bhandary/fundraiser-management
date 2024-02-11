package repository

const insertDonationQuery = `
	insert into donation (user_id, fundraiser_id, amount, is_anonymous) 
	values($1, $2, $3, $4) Returning id;
`

const listUserDonationsQuery = `
	select 
		donation.id, donation.fundraiser_id, fundraiser.title as fundraiser_title, 
		donation.amount, donation.is_anonymous, donation.created_at
	from donation
	join fundraiser
	on donation.fundraiser_id = fundraiser.id
	where	(donation.user_id = $1) and 
		(CASE when $7 in ('true', 'false') THEN is_anonymous = ($7 = 'true') ELSE true END) and
		(fundraiser.title ilike '%' || $2 || '%')
	order by 
	    CASE WHEN $3 = 'donation_id' and $4 THEN donation.id END ASC,
	    CASE WHEN $3 = 'donation_id' and not $4 THEN donation.id END DESC,

	    CASE WHEN $3 = 'fundraiser_id' and $4 THEN donation.fundraiser_id END ASC,
	    CASE WHEN $3 = 'fundraiser_id' and not $4 THEN donation.fundraiser_id END DESC,

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

const getListUserDonationsCountQuery = `
	select count(donation.id)
	from donation
	join fundraiser
	on donation.fundraiser_id = fundraiser.id
	where	(donation.user_id = $1) and 
		(CASE when $3 in ('true', 'false') THEN is_anonymous = ($3 = 'true') ELSE true END) and
		(fundraiser.title ilike '%' || $2 || '%');
`

const listFundraiserDonationsQuery = `
	select 
		donation.id, 
		(CASE WHEN donation.is_anonymous then 0 else donation.user_id END) as user_id, 
		(CASE WHEN donation.is_anonymous then 'Hidden' else users.first_name || ' ' || users.last_name END) as user_name,
		donation.fundraiser_id, fundraiser.title as fundraiser_title, 
		donation.amount, donation.is_anonymous, donation.created_at
	from donation
	join fundraiser
	on donation.fundraiser_id = fundraiser.id
	join users
	on donation.user_id = users.id
	where donation.fundraiser_id = $1
	order by donation.created_at DESC
	offset $2 
	limit $3;
`

const getListFundraiserDonationsCountQuery = `
	select count(donation.id) 
	from donation
	where donation.fundraiser_id = $1;
`

const listDonationsQuery = `
	select 
		donation.id, donation.user_id, (users.first_name || ' ' || users.last_name) as user_name,
		donation.fundraiser_id, fundraiser.title as fundraiser_title, 
		donation.amount, donation.is_anonymous, donation.created_at
	from donation
	join fundraiser
	on donation.fundraiser_id = fundraiser.id
	join users
	on donation.user_id = users.id
	where	(
			users.first_name ilike '%' || $1 || '%' or 
			users.last_name ilike '%' || $1 || '%' or
			fundraiser.title ilike '%' || $1 || '%' 
		) and	
		(CASE when $2 in ('true', 'false') THEN is_anonymous = ($2 = 'true') ELSE true END)
	order by 
	    CASE WHEN $3 = 'donation_id' and $4 THEN donation.id END ASC,
	    CASE WHEN $3 = 'donation_id' and not $4 THEN donation.id END DESC,

	    CASE WHEN $3 = 'fundraiser_id' and $4 THEN donation.fundraiser_id END ASC,
	    CASE WHEN $3 = 'fundraiser_id' and not $4 THEN donation.fundraiser_id END DESC,

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

const getListDonationsCountQuery = `
	select count(donation.id)
	from donation
	join fundraiser
	on donation.fundraiser_id = fundraiser.id
	join users
	on donation.user_id = users.id
	where	(
			users.first_name ilike '%' || $1 || '%' or 
			users.last_name ilike '%' || $1 || '%' or
			fundraiser.title ilike '%' || $1 || '%' 
		) and	
		(
			CASE when $2 in ('true', 'false') THEN is_anonymous = ($2 = 'true') ELSE true END
		);
`
