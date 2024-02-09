package repository

const insertDonationQuery = `
	insert into donation (user_id, fundraiser_id, amount, is_anonymous) 
	values($1, $2, $3, $4) Returning id;
`

const listUserDonations = `
	select 
		donation.id, donation.fundraiser_id, fundraiser.title as fundraiser_title, 
		donation.amount, donation.is_anonymous, donation.created_at
	from donation
	join fundraiser
	on donation.fundraiser_id = fundraiser.id
	where donation.user_id = $1
	order by created_at desc;
`
