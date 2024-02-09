package repository

const insertDonationQuery = `
	insert into donation (user_id, fundraiser_id, amount, is_anonymous) 
	values($1, $2, $3, $4) Returning id;
`
