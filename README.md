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
git clone https://github.com/Unpatches/pz10-auth
cd pz9-auth
```

Команда запуска сервера:

```
go run ./cmd/server
```
```
make docker-build
```
```
make docker-run
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
### Логин админа POST
```
http://185.250.46.179:8082/api/v1/login
```
```
{"email":"admin@example.com","password":"secret123"}
```

<img width="825" height="495" alt="Снимок экрана 2025-12-09 в 17 21 29" src="https://github.com/user-attachments/assets/4e57d47f-ca87-4745-aeb0-75e5388b0ba5" />


### Логин обычного пользователя (user) POST
```
http://185.250.46.179:8082/api/v1/login
```
```
{"email":"user@example.com","password":"secret123"}
```

<img width="817" height="507" alt="Снимок экрана 2025-12-09 в 17 43 50" src="https://github.com/user-attachments/assets/ef570f56-e880-4414-b952-90d348ec25cc" />



### Проверка /api/v1/me
```
http://185.250.46.179:8082/api/v1/me
```

<img width="634" height="418" alt="Снимок экрана 2025-12-09 в 17 56 15" src="https://github.com/user-attachments/assets/316a63e1-d9ad-498e-9db0-a0e0d7dc2b4d" />



### Проверка /api/v1/admin/stats
```
http://185.250.46.179:8082/api/v1/admin/stats
```

<img width="606" height="377" alt="Снимок экрана 2025-12-09 в 18 04 54" src="https://github.com/user-attachments/assets/ace420ad-23be-4fee-bb71-3f970cf894c5" />


------
