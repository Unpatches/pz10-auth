## Имя: Дорджиев Виктор
## Группа: ЭФМО-02-25
# Проект pz10-auth

Задачи проекта
- Понять устройство JWT и где его уместно применять в REST API. 
- Сгенерировать и проверить JWT в Go (HS256), передавать его в Authorization: Bearer …. 
- Реализовать middleware-аутентификацию (достаёт токен, валидирует, кладёт клеймы в context). 
- Добавить middleware-авторизацию (RBAC/права на эндпоинты). 
- Встроить это в уже знакомую архитектуру HTTP-сервиса/роутера.


---

## Установка и запуск

(Необходимы предустановленные Go версии 1.25 и выше и Git)

Клонировать репозиторий:

```
git clone https://github.com/Unpatches/pz9-auth
cd pz9-auth
```

Команда запуска сервера:

```
go run ./cmd/server
```


------

## Структура проекта

```plaintext
pz10-auth/
├── go.mod
├── go.sum
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── core/
│   │   ├── user.go
│   │   └── service.go
│   ├── http/
│   │   ├── router.go
│   │   └── middleware/
│   │       ├── authn.go
│   │       └── authz.go
│   ├── platform/
│   │   ├── config/
│   │   │   └── config.go
│   │   └── jwt/
│   │       └── jwt.go
│   └── repo/
│       ├── user_mem.go
│       └── refresh_mem.go
├── README.md
├── go.mod
├── go.sum
└── .gitignore       
```

## Отчёт о проделанной работе
### Регистрация POST
```
http://185.250.46.179:8081/auth/register
```
```
{"email":"user@example.com","password":"Secret123!"}
```

<img width="457" height="505" alt="image" src="https://github.com/user-attachments/assets/33face01-2e6b-4f2d-96fa-74a6bc5b52cd" />


### Повторная регистрация POST
```
http://185.250.46.179:8081/auth/register
```
```
{"email":"user@example.com","password":"AnotherPass"}
```

<img width="454" height="447" alt="image" src="https://github.com/user-attachments/assets/d48c407c-6581-408d-a3c8-3c58ffb33bbf" />


### Вход (верный) POST
```
http://185.250.46.179:8081/auth/login
```
```
{"email":"user@example.com","password":"Secret123!"}
```

<img width="409" height="501" alt="image" src="https://github.com/user-attachments/assets/ed2d4711-fbe1-4731-8446-76ddc766a1cd" />



### Вход (неверный) POST
```
http://185.250.46.179:8081/auth/login
```
```
{"email":"user@example.com","password":"wrong"}
```

<img width="417" height="445" alt="image" src="https://github.com/user-attachments/assets/85c6ef6f-d9b2-4596-b0a5-84b31ddeca63" />

------## Имя: Дорджиев Виктор
## Группа: ЭФМО-02-25
# Проект pz9-auth

Задачи проекта
- Научиться безопасно хранить пароли (bcrypt), валидировать вход и обрабатывать ошибки.
- Реализовать эндпоинты POST /auth/register и POST /auth/login.
- Закрепить работу с БД (PostgreSQL + GORM или database/sql) и валидацией ввода.
- Подготовить основу для JWT-аутентификации в следующем ПЗ


---

## Установка и запуск

(Необходимы предустановленные Go версии 1.25 и выше и Git)

Клонировать репозиторий:

```
git clone https://github.com/Unpatches/pz9-auth
cd pz9-auth
```

Команда запуска сервера:

```
go run ./cmd/api
```


------

## Структура проекта

```plaintext
pz9-auth/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── core/
│   │   └── user.go
│   ├── http/
│   │   └── handlers/
│   │       └── auth.go
│   ├── platform/
│   │   └── config/
│   │       └── config.go
│   └── repo/
│       ├── postgres.go
│       └── user_repo.go
├── go.mod
├── go.sum
└── .gitignore       
```

## Отчёт о проделанной работе
### Регистрация POST
```
http://185.250.46.179:8081/auth/register
```
```
{"email":"user@example.com","password":"Secret123!"}
```

<img width="457" height="505" alt="image" src="https://github.com/user-attachments/assets/33face01-2e6b-4f2d-96fa-74a6bc5b52cd" />


### Повторная регистрация POST
```
http://185.250.46.179:8081/auth/register
```
```
{"email":"user@example.com","password":"AnotherPass"}
```

<img width="454" height="447" alt="image" src="https://github.com/user-attachments/assets/d48c407c-6581-408d-a3c8-3c58ffb33bbf" />


### Вход (верный) POST
```
http://185.250.46.179:8081/auth/login
```
```
{"email":"user@example.com","password":"Secret123!"}
```

<img width="409" height="501" alt="image" src="https://github.com/user-attachments/assets/ed2d4711-fbe1-4731-8446-76ddc766a1cd" />



### Вход (неверный) POST
```
http://185.250.46.179:8081/auth/login
```
```
{"email":"user@example.com","password":"wrong"}
```

<img width="417" height="445" alt="image" src="https://github.com/user-attachments/assets/85c6ef6f-d9b2-4596-b0a5-84b31ddeca63" />

------
