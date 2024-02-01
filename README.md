# Product App Go Project
## Run Setup

#### Follow the steps below in order to run the project in local:

To get the frameworks used in the project:
- echo framework - https://github.com/labstack/echo (for webserver)
  - go get github.com/labstack/echo/v4
- pgxpool - https://github.com/jackc/pgx (for database)
  - go get github.com/jackc/pgx/v4/pgxpool@v4.18.1
- gommon - https://github.com/labstack/gommon (for logging)
  - go get github.com/labstack/gommon
- testify - https://github.com/stretchr/testify (for testing)
  - go get github.com/stretchr/testify


#### To create a test environment
- Create a file with the name of **test_db.sh** under the test/scripts directory. Add the following line in it.
> If you are using **Windows** as an operating system, you must install the **git bash(https://git-scm.com/download)** tool before you can execute the sh file. If you are using **Linux** or **Mac OS X**, you can run the sh file directly.

```bash
#!/bin/bash
```
- Add the following lines inside the **test_db.sh** file to pull the postgres image from DockerHub.
```bash
docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 6432:5432 -d postgres:latest
echo "Postgresql starting..."
sleep 3
```
- Add the following lines inside the **test_db.sh** to execute the create database psql command in it.
```bash
docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE productapp"
sleep 3
echo "Database productapp created"
```
- Add the following lines inside the **test_db.sh** to execute the create table psql command in it.
```bash
docker exec -it postgres-test psql -U postgres -d productapp -c "create table if not exists products
(
  id bigserial not null primary key,
  name varchar(255) not null,
  price double precision not null,
  discount double precision,
  store varchar(255) not null
);
"
sleep 3
echo "Table products created"
```
- Go to the scripts folder in terminal and execute the following command.
```bash
sh test/scripts/test_db.sh
```
- See the docker container is running...

![Postgres instance running](/test/documentation/docker-postgres-running.png "Postgres instance running")

- Add data sources and drivers

![Postgres db config](/test/documentation/postgres-db-config.png "Postgres db config")

- See the postgres table and its columns

![Postgres db table](/test/documentation/postgres-db-table.png "Postgres db table")


- You can run all the tests with the command below.
```bash
go test test/infrastructure/product_repository_test.go -v
```