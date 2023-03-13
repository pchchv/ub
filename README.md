# ub - user balance

REST API system for registering users and storing their balances

### Running the application

```
go run .
```

### Running tests (app must be running)

```
go test .
```

## HTTP Methods

```
"GET" / — Checking the server connection

    example: 
        "GET" :8080/
```

```
"GET" /ping — Checking the server connection

    example: 
        "GET" :8080/ping
```
```
"GET" /user — Get a user data

    options:
        id — User ID

    example: 
        "GET" :8080/user?id=sadf54-fdsa48-dsaf459-dsaf45
```

```
"POST" /user — Create a new user. Need JSON body

    example: 
        "POST" :8080/create
```

```json
{
    "email" : "ipchchv@gmail.com",
	"name" : "Jack",
	"password" : "3223414r"
}
```

```
"PATCH" /balance — Update user balance. Need JSON body
```

```json
{
    "id" : "sadf54-fdsa48-dsaf459-dsaf45",
	"operation" : "deposit",
	"amount" : 175.5
}
```

```
"DELETE" /user — Delete one User

    options:
        id — User ID

    example: 
        "DELETE" :8080/user?id=sadf54-fdsa48-dsaf459-dsaf45
```
