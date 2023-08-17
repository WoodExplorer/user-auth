# user-auth

A simple implementation of user auth manager.

Hopefully, this implementation could provide relevant information about what my skills and habit.

[![Go Report Card](https://goreportcard.com/report/github.com/WoodExplorer/user-auth)](https://goreportcard.com/report/github.com/WoodExplorer/user-auth)

## Usages

### prerequisites

development tools：
- golang 1.20
- (optional) docker, if you would like to build docker images
- (optional) make, if you would like to use it

### docker images

Assuming related development tools are installed, you can use `make builder` or you can execute following commands, :
```bash 
docker build -t {your_image_name} -f build/Dockerfile .
```

### test coverage

Assuming related development tools are installed, you can use `make cover` or you can execute following commands:
```bash
go test ./... -cover -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### curl

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

## Thoughts

Despite the requirements, ambiguities do exist. I tried to handle them properly. But what I did might not be the same as you might expect. If this kind of unfortunate situation happens, I would like to suggest that more communication can improve it.   

### Key problems and my solutions

This assignment seems simple, but it could reflect a candidate's ability. 

After reading the requirements, several problems came to me:
1. If data is not persistent, then data should be stored in memory. What data structure should we use?
2. To what extent should one handle the problems regarding concurrent requests?

To 1st problem, I look up to redis. I tried to implement a redis-like memory store, providing basic SET-GET-DEL operations and HSET-HGET-HDELALL operations. Such an arrangement can be handy, but it can cause problems related to entity relationships.

The 2nd problem mainly refers to the data level. I mean, implementing a fledged MVCC style is difficult. So, I had to choose something more easily implemented. In implementing the redis-like memory store, not only did I imitate the data types, but also the 1-thread for data handling design. Specifically, a single goroutine will handle the data operations while requests are sent to it via a channel. Also, to provide proper function, transaction must be supported so as to avoid problems that arise when isolation is lacking. A preliminary tx is provided, but I understand it is far from completion.    


### API

As I understand, APIs provided here are Restful, and its implementation follows a layered solution: router->service->repository->data storage.  

### Performance

Here is a too-short complexity analysis.

#### time

Basically using map, time complexity is approximately O(1).

#### space

Basically using map, space complexity is approximately O(n). 

### Testing

Regretfully, test coverage is in need of improvement due to time limit. I always appreciate the importance of testing, and I am willing to improve it.  

## Known issues
- when delete a role, corresponding user-role records has not been deleted yet
- due to time limit, test coverage is in need of improvement
- due to time limit, testing code can be optimized to a larger extent
- has not use wire to manage object dependency yet
- no error code yet except for 0 - ok and 1 - error 
- no error message i18n
- no id field for user and role yet, due to lack of a proper id generation scheme, which I found difficult to implement in a short time window.
- ...

## Others
 