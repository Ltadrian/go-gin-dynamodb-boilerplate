Rewrite of https://github.com/vsouza/go-gin-boilerplate

Updated libraries, easier to set up dynamodb with init script for local development

Download modules:
go mod download

## Development setup

Running DynamoDB on Docker Image:
https://hub.docker.com/r/amazon/dynamodb-local/ <br/>

docker pull amazon/dynamodb-local

Set up database:
./db/init_database.sh

To Run: <br/>
go run .

## JWT Authentication

1. JWT tokens can be generated from the GET /token endpoint
2. JWT tokens can be validated in the middleware via auth.go
3. JWT tokens need a secret for signing and verifying, and personally what I recommend for production is NOT by putting it in source code, but rather pass in the randomly generated secret via environmental variables. I show an example of environmental variable through configuration


## Usage

# Health Endpoint
GET /health

# Token Endpoint
GET /v1/token

# User Endpoints
POST /v1/user/
Body {
    "Name": "Joe",
    "Gender": "Male",
    "Age": 25
}

GET /v1/user/9619e438-ffa5-4a7f-a160-e2203593d5eb

## TODO
<li> tests </li>
<li> logs </li>
<li> makefile </li>
<li> docker image </li>
<li> auto generated swagger documentation </li>
