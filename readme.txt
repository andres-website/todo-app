## Как запустить контейнер с Postgres в Docker

docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres 


## Миграция Postgres

migrate create -ext sql -dir ./schema -seq init 


## Но для начала мигрций надо установить scoop в Windows

Set-ExecutionPolicy RemoteSigned -Scope CurrentUser # Optional: Needed to run a remote script the first time
irm get.scoop.sh | iex

### А затем установить migrate

scoop install migrate


### Применяем миграцию к БД
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up


### Как попасть к базе данных внутрь докера

docker exec -it 6fc10df37be9 /bin/bash
// 6fc10df37be9 - id контейнера который можно узнать по docker ps

psql -U postgres
\d


### Что делать, если накосячил с миграциями

// Внутри psql
update schema_migrations set version='000001', dirty=false;

select * from schema_migrations;