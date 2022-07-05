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

## TODO
<li> jwt authentication </li>
<li> middleware </li>
<li> logs </li>
<li> makefile </li>
<li> docker image </li>
