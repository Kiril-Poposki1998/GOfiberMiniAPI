# Express based API
## Build the application with Docker
```
docker build -t fiber_app .
```
## Run the application
```
docker run -p 8080:8080 --rm fiber_app 
```
### Types of HTTP requests to paths
```
GET /api/person - displays every person
GET /api/person/id - displays the person with the id
POST /api/person/id - add the person
PUT /api/person/id - update the person
DELETE /api/person/id - delete the person
POST /auth/register - register a new user
POST /auth/login - login with a user and password
```
Example login of a user (the same goes for registering):
```
{
    "username": "user",
    "password: "password"
}
```
Example of creating/editing a person:
```
{
    "name": "John",
    "surname": "Doe"
}
```