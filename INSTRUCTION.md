# User Service

## Known Issue with Go 1.18

If you do local build in Mac, you may encountered an error such `//go:linkname must refer to declared function or variable`
please do the following command `go get -u golang.org/x/sys`

## Build and start the app using Docker Compose

```bash
docker compose up --build
```

The docker compose contains both User Service and postgres database.

## Testing API with cURL

### Signup

#### Signup Request

```bash
curl --location --request POST 'http://localhost:8088/usermanagement/v1/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstName": "John",
    "lastName": "Doe",
    "email": "John.Doe@jdoe.org",
    "password": "pwd"
}'
```

#### Signup Response

```bash
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkpvaG4uRG9lQGpkb2Uub3JnIiwiZXhwaXJhdGlvbiI6MTY1MTk2NzY4Mn0.bgvvDqmG2NolXK7Sf7GEVyMdVjZ3CG81RBgSxshi100}
 ```

### Login

#### Login Request

 ```bash
curl --location --request POST 'http://localhost:8088/usermanagement/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "John.Doe@jdoe.org",
    "password": "pwd"
}'
```

#### Login Response

```bash
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkpvaG4uRG9lQGpkb2Uub3JnIiwiZXhwaXJhdGlvbiI6MTY1MTk2NzcwMH0.ItrIMT9koidG2k_pgrKH8uaEaACEfBd_h180c-fRICI"}
 ```

### Update

#### UpdateRequest

Note: Please use the token instead of copy from below that's either genrated from Signup or Login, as the token has claims for  email and expiration, and the system will validate it  

The following request will update firstName and lastName: John --> John2 and Doe --> Doe2

```bash
curl --location --request PUT 'http://localhost:8088/usermanagement/v1/users' \
--header 'x-authentication-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkpvaG4uRG9lQGpkb2Uub3JnIiwiZXhwaXJhdGlvbiI6MTY1MTk2NzcwMH0.ItrIMT9koidG2k_pgrKH8uaEaACEfBd_h180c-fRICI' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstName": "John2",
    "lastName": "Doe2",
    "email": "John.Doe@jdoe.org",
    "password": "pwd"
}'
```

#### UpdateResponse

```bash
Empty
```

### List users

#### ListRequest

Note: Please use the token instead of copy from below that's either genrated from Signup or Login, as the token has claims for  email and expiration, and the system will validate it

```bash
curl --location --request GET 'http://localhost:8088/usermanagement/v1/users' \
--header 'x-authentication-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IkpvaG4uRG9lQGpkb2Uub3JnIiwiZXhwaXJhdGlvbiI6MTY1MTk2NzcwMH0.ItrIMT9koidG2k_pgrKH8uaEaACEfBd_h180c-fRICI' \
--header 'Content-Type: application/json'
```

#### ListResponse

Note: The response reflected that FirstName and LastName have been updated.

```bash
{"users":[{"email":"John.Doe@jdoe.org","firstName":"John2","lastName":"Doe2"}]}
```
