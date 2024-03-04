
# Fundraiser Management System in Go

## Use Case

### There are 3 types of actors
    
    User (Donors)
    Organizer (organizes fundraisers)
    Admin

1. Donor

- Register and Login on the platform
- View fundraisers by different Organizers
- Search and filter through fundraisers
- View Detail of an specific fundraiser
- View list of donors who have donated to a fundraiser (Sorted in amount of donations)
- Donate to a fundraiser
- Keep their donation anonymous

2. Organizer
- Request to register on the platform
- Create and Edit fundraiser
- View list of users who have donated to a specific fundraiser

3. Admin
- Monitor everything
- Verify organizer and grant access
- Ban a fundraiser
- Edit and delete an organize


## Prerequisites

- go
- Postgresql
- Make

## Setup

Clone the project

```bash
  git clone https://github.com/Suhaan-Bhandary/fundraiser-management.git
```

Open Postgresql in terminal

```bash
sudo -u postgres -i 
psql
CREATE DATABASE fundraiser_management;
\c fundraiser_management # changing the database
```

Once the database is created go to the project and copy the up file from db/migrations folder and paste it in psql to setup tables

Setup Env in root directory, refer the Environment Variables section

```bash
  cd fundraiser-management
``` 

Start the server

```bash
  make run
```


## Swagger

Swagger file is provided with the project to view routes and their working

## Postman Collection

[Postman Collection](https://www.postman.com/mission-architect-94960085/workspace/public/collection/16036286-a6c28b2e-5b8f-4267-b6a7-e999b9e02e7e?action=share&creator=16036286)
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`PSQL_INFO`

`SECRET_KEY`

Note:
PSQL_INFO format = "host=<> port=<> user=<> password=<> dbname=<> sslmode=disable"


## Running Tests

To run tests, run the following command

### Test

```bash
  make test
```

### Test with coverage Detail 
```bash
  make test-cover
```

### Test with coverage Detail in browser
```bash
  make html-cover
```
## Authors

- [@suhaanbhandary](https://suhaan-bhandary.vercel.app/)

