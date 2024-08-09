Introduction: 
This project is collaborative voting system backend that allows multiple users to create,
join, and participate in voting sessions. The application should demonstrate the use of WebSockets for real-
time updates, Go-routines for handling concurrent connections, and routes for managing voting operations and
user authentication.

Getting Started: Prerequisites: Golang (version 1.22.x)

Installation: 
1. Clone the repository: gh repo clone GautamViki/CollaborativeVotingSystem 
2. Install dependencies: go mod tidy

Running the Application: go run main.go


Features: 
1. Register a user with username and password
2. Generate a jwt token and validate the token 
3. Cast there vote

**********User Registration*************
Request Body:
Endpoint =>  POST: http://localhost:3009/user/register
curl --location 'http://localhost:3009/user/register' \
--header 'Content-Type: application/json' \
--data '{
    "username":"vikasfas",
    "password":"123safdf"
}'

Response Body:
   User Registered.


*****************Token Genration****************
Request body:
Endpoint => POST: http://localhost:3009/token
curl --location 'http://localhost:3009/token' \
--header 'Content-Type: application/json' \
--data '{
    "username":"vikasfas",
    "password":"123safdf"
}'

Response Body:
{
    "token_type": "Bearer",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMzMDMwMDUsInBhc3N3b3JkIjoiMTIzc2FmZGYiLCJ1c2VybmFtZSI6InZpa2FzZmFzIn0.M9Fc0PDDD6b_GCWrz0LtinXjQzOgfWEPfWYBRMeGHJE",
    "expires_in": 1723216665
}

**********************Cast the Vote**************************
Request Body: 
Endpoint => POST : http://localhost:3009/user/castvote
curl --location 'http://localhost:3009/user/castvote' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMzMDIwNTcsInBhc3N3b3JkIjoiMTIzc2FmZGYiLCJ1c2VybmFtZSI6InZpa2FzZmFzIn0.Bs5g7DuyLb8FG2iIEDKX0Z-9N5JuAK8Pp_m2Duhn4Cw' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMzMDIwNTcsInBhc3N3b3JkIjoiMTIzc2FmZGYiLCJ1c2VybmFtZSI6InZpa2FzZmFzIn0.Bs5g7DuyLb8FG2iIEDKX0Z-9N5JuAK8Pp_m2Duhn4Cw' \
--data '{
    "username":"vikasfas",
    "password":"123safdf"
}'

Response Body:
{"username":"vikasfas","password":"123safdf"}
