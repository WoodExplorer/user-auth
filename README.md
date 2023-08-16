# user-auth

## coverage

```bash
go test ./... -cover -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## thoughts
atomicity
isolation

## time complexity analysis

### known issues
- no error code yet except for 0 - ok and 1 - error 
- no error message i18n
- no id field for user and role yet, due to lack of a proper id generation scheme, which I found difficult to implement in a short time window.
- has not use wire to manage object dependency yet

### curl usages

- create user
```bash
curl -L -X POST http://127.0.0.1:8080/api/v1/users -H 'Content-Type: application/json' -d '{"name":"ua","password":"pwd"}'
```

- list users
```bash
curl http://127.0.0.1:8080/api/v1/users 
```

- create role
```bash
curl -L -X POST http://127.0.0.1:8080/api/v1/roles -H 'Content-Type: application/json' -d '{"name":"ra"}'
```

- list roles
```bash
curl http://127.0.0.1:8080/api/v1/roles 
```

- bind 
```bash
curl -L -X POST http://127.0.0.1:8080/api/v1/user-roles -H 'Content-Type: application/json' -d '{"userName":"ua","roleName":"ra"}'
```

- authenticate
```bash
 curl -L -X POST http://127.0.0.1:8080/api/v1/authn/tokens  -H 'Content-Type: application/json' -d '{"name":"ua","password":"pwd"}'
```

- invalidate
```bash
 curl -L -X DELETE http://127.0.0.1:8080/api/v1/authn/tokens  -H 'Content-Type: application/json' -d '{"token":"{token}"}'
```

- check-role
```bash
 curl  'http://127.0.0.1:8080/api/v1/authz/check-role?roleName=ra&token={token}'
```

- user-roles
```bash
 curl  'http://127.0.0.1:8080/api/v1/authz/user-roles?token={token}'
```
