## Аля make

docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

go run .\cmd\main.go


## Аля Тесты мануальные

1. Зарегестрировать пользователя (sign-up)
POST localhost:8008/auth/sign-up
{
    "name": "Andres",
    "username": "firstAndres",
    "password": "qwerty"
}
// assert: возвращает id пользователя


2. Получить JWT токен (авторизоваться) (sign-in)
POST localhost:8008/auth/sign-in
{
    "username": "firstAndres",
    "password": "qwerty"
}
// assert: возвращает token пользователя


3. Попытаться доступиться до списков тудушек (c Bearer authorization) 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDIxNTIzNTcsImlhdCI6MTcwMjEwOTE1NywidXNlcl9pZCI6MX0.5U7Y3x0cfqfFFeNIuF-VFrZypbmNnEtnPP2WaoP2BlA
GET http://localhost:8008/api/lists
bearer <token>
assert: возвращает  в поле data данные или null - если данных нет


4. Создаёт список тудушек
POST http://localhost:8008/api/lists
{
    "title": "Список важных дел",
    "description": "Срочно нужно сделать"
}
assert: возвращает  id созданного списка


5. Возвращает список тудушек по id списка
GET http://localhost:8008/api/lists/1
assert: возвращает структуру списка (id, title, description)


6. Редактирует список тудушек по id списка
PUT http://localhost:8008/api/lists/1
{
    "title": "Список 4",
    "description": "Описание Список 4"
}
assert: возвращает status ok


7. Создаёт Тудушку по id списка (и id текущего пользователя из контекста, по JWT токену)
POST https://youtu.be/zpMHy4UVDro?si=kuud2a_9HRf9THZ5&t=250
{
    "title": "Арбуз"
}
assert: возвращает id тудушки


8. Возрвращает все тудушки по id списка (и id текущего пользователя из контекста, по JWT токену)
GET http://localhost:8008/api/lists/1/items
assert: возвращает все тудушки из списка



## Как запустить контейнер с Postgres в Docker

docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres 
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres


## Original

https://github.com/zhashkevych/todo-app


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

docker exec -it 3dd853822b96 /bin/bash
// 6fc10df37be9 - id контейнера который можно узнать по docker ps

psql -U postgres
\d
select * from users;


#### Вариант 2: (очень быстрый) Как попасть к базе данных внутрь докера
docker exec -it todo-db psql -U postgres -d postgres


### Что делать, если накосячил с миграциями

// Внутри psql
update schema_migrations set version='000001', dirty=false;

select * from schema_migrations;


## Установка Viper
go get -u github.com/spf13/viper

## Установка SQL X
go get -u github.com/jmoiron/sqlx

## Установка go dot env
go get -u github.com/joho/godotenv

## Установка logrus
go get -u github.com/sirupsen/logrus

## Установка JWT
go get -u github.com/dgrijalva/jwt-go


## Что бы завелась реализация драйвера для Postgres
go mod download github.com/lib/pq


# Запуск get запроса (в хендлере которого лежит tele_noti("..."))
http://localhost:8008/api/lists

https://www.youtube.com/watch?v=QTFoGgLqTYA


## Аля make

docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

go run .\cmd\main.go


## Запросы для POSTman

*Получение списка листов*
http://localhost:8008/api/lists

*Регистрация*
http://localhost:8008/auth/sign-up

{
    "name": "Andres",
    "username": "firstAndres",
    "password": "qwerty"
}


### Ошибка: error: "socket hang up" в Postman

Причина крылась в файле pkg/repository/postgres.go
Функция конструктор: NewPostgresDB (...)дщ                                                                                                                                                                                                                                                                                                                                        

В итоге правильно так: return db, err


## Next lavel
https://youtu.be/zpMHy4UVDro?si=kuud2a_9HRf9THZ5&t=250