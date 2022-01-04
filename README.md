# A basic project to learn how bare basic docker works

## Application Overview
* Fetch data from coingecko API using golang server and store in mongoDB.
* Serve simple Web frontend using nginx.
* Add coins in market watch and wait for price updates.

Server, frontend and DB are hosted inside an docker container.

Install docker and just run `docker-compose -f docker-compose.yaml up`
