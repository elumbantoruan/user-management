# User Management
## API Specs

### `POST /signup`
Endpoint to create an user row in postgres db. The payload should have the following fields:

```json
{
  "email": "edison@pellucidcomputing.com",
  "password": "elumbantoruan",
  "firstName": "Edison",
  "lastName": "Lumbantoruan"
}
```

where `email` is an unique key in the database.

The response body should return a JWT on success that can be used for other endpoints:

```json
{
  "token": "some_jwt_token" 
}
```

### `POST /login`
Endpoint to log an user in. The payload should have the following fields:

```json
{
  "email": "edison@pellucidcomputing.com",
  "password": "mysecret"
}
```

The response body should return a JWT on success that can be used for other endpoints:

```json
{
  "token": "some_jwt_token"
}
```

### `GET /users`
Endpoint to retrieve a json of all users. This endpoint requires a valid `x-authentication-token` header to be passed in with the request.

The response body should look like:
```json
{
  "users": [
    {
      "email": "edison@pellucidcomputing.com",
      "firstName": "Edison",
      "lastName": "Lumbantoruan"
    }
  ]
}
```

### `PUT /users`
Endpoint to update the current user `firstName` or `lastName` only. This endpoint requires a valid `x-authentication-token` header to be passed in and it should only update the user of the JWT being passed in. The payload can have the following fields:

```json
{
  "firstName": "NewFirstName",
  "lastName": "NewLastName"
}
```

The response can body can be empty.


