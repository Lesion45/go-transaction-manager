# Transaction manager
This project is an API for working with user balances. [Here](https://github.com/avito-tech/internship_backend_2022) you can see a TA.
## Features
* Crediting of funds.
* Writing of funds.
* Retrieving user balance.
* Getting reports (Not released yet).
## Configuration
The API and storage settings are stored in `config.yaml`. 
The structure of this file should be as follows:
```
env: "prod" # options: local dev prod

storage:
  host: "db"
  port: 5432
  user: "username"
  password: "your.password"
  dbname: "your.storage.name"

http-server:
  address: "0.0.0.0:5000"
  timeout: 10s
  idle-timeout: 60s
```
You also have to update `docker-compose.yml` after changing `config.yml`:
```
version: '3.8'

services:
  app:
    container_name: transaction-manager
    build: ./
    ports:
      - 5000:5000
    restart: on-failure
    depends_on:
      - db
    networks:
      - your.network

  db:
    image: postgres:latest
    container_name: db_postgres
    environment:
      - POSTGRES_USER=username
      - POSTGRES_PASSWORD=your.password
      - POSTGRES_DB=your.storage.name
      - POSTGRES_HOST=localhost
    volumes:
      - ./sql:/docker-entrypoint-initdb.d
    networks:
      - your.network

networks:
  your.network:
    driver: bridge
```
## Installation
1. Clone the repository:
```
git clone https://github.com/Lesion45/go-transaction-manager.git
cd go-transaction-manager
```
2. Run makefile:
```
make start
```
## Usage
### Crediting of funds
Endpoint: `api/v1/user/add_balance`, Method: `POST`, Content-type: `application/json`
#### Request
```
curl --location --request POST 'http://localhost:5000/api/v1/user/add_balance' \
--header 'Content-Type: application/json' \
--data-raw '{ \
    "amount": 1000.0,
    "user-id": "2c220c36-6d64-40ae-9ae2-8b8631208d4d"
}'
```
#### Response
```
{
    "status": "Ok"
}
```
### Reserving of founds
Endpoint: `api/v1/reservation/reserve_balance`, Method: `POST`, Content-type: `application/json`
#### Request
```
curl --location --request POST 'http://localhost:5000/api/v1/reservation/reserve_balance' \
--header 'Content-Type: application/json' \
--data-raw '{ \
    "amount": 1000.0,
    "info": PlayStation 5,
    "order-id": "91b0e054-0970-4a24-8a9e-d8a592cfbd44",
    "service-id": "fa9b7218-dba3-4b94-88cf-c5b69f4e8d24",
    "user-id": "2c220c36-6d64-40ae-9ae2-8b8631208d4d"
}'
```
#### Response
```
{
    "status": "Ok"
}
```
### Committing of reserved founds
Endpoint: `api/v1/reservation/commit_balance`, Method: `POST`, Content-type: `application/json`
#### Request
```
curl --location --request POST 'http://localhost:5000/api/v1/reservation/commit_balance' \
--header 'Content-Type: application/json' \
--data-raw '{ \
    "amount": 1000.0,
    "info": PlayStation 5,
    "order-id": "91b0e054-0970-4a24-8a9e-d8a592cfbd44",
    "service-id": "fa9b7218-dba3-4b94-88cf-c5b69f4e8d24",
    "user-id": "2c220c36-6d64-40ae-9ae2-8b8631208d4d"
}'
```
#### Response
```
{
    "status": "Ok"
}
```
### Retrieving user balance
Endpoint: `api/v1/user/get_balance`, Method: `GET`, Content-type: `application/json`
#### Request
```
curl --location --request GET 'http://localhost:5000/api/v1/user/get_balance' \
--header 'Content-Type: application/json' \
--data-raw '{ \
    "user-id": "2c220c36-6d64-40ae-9ae2-8b8631208d4d"
}'
```
#### Response
```
{
    "balance": 1000.0
}
```
## API Documentation
API documentation is available at:

Swagger UI: http://localhost/swagger
## TODO
* Clean the code
* Add middleware for collecting metrics
* Increase test coverage
## Author
* [Lesion45](https://github.com/Lesion45)