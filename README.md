## Hello there!

You have to implement a REST API for a simple user management system.

## Requirements / user stories
Open ./api/api.yaml in swagger viewer eg.: editor.swagger.io... Then you see 4 rest endpoints. You need implement a Post,Put and Get(listing) endpoint.
I was implement an example endpoint -» user/count Please follow this pattern, if you can.


You don't need:
- Api test case
- Custom error handling, just internal server error

## About the project struct

## API development

The `api` folder contains the Open API 3 Specification for the service endpoints. Using the `make codegen` command,
the developers can easily extend the REST API. By running the `make codegen` script, we generate the routes by extending
the `ServerInterface` method list with the new endpoints. Our `HTTPServer` structure implements the `ServerInterface`.

## Code structure

All external inputs hits the `ports` package. The only entry point to the application is through the `Ports layer`
(HTTP handlers, Pub/Sub message handlers). `Ports` execute relevant handlers in the `App` layer. Some of these will call
the `Domain` code, and some will use `Adapters`, which are the only way out of the service. The `adapters` layer is where
your database queries and HTTP clients live.
![](resources/ddd_base_arch.jpeg)


### App.env

It is mandatory to have the `app.env` file in your root folder. We added a `app.env.example` file to be able to copy-paste-rename it to quickly spin up the project.

### Configuration

`config/config.go` file contains all the available configuration variables.
The config loader loads the values from the `app.env` file. The loader checks for overrides in the
environment variables. Config loader uses the following precedence order. Each item takes precedence over the item below it:

* explicit call to Set
* flag
* env
* config
* key/value store
* default

Used library: [viper](https://github.com/spf13/viper)

# Installation

The service uses `Go 1.18`

### Download dependencies

```bash
$ go mod download
``` 

### Start PostgreSQL

```bash
# if you want to start a new postgresql server within docker
$ make docker-up

# if you want to stop & remove the postgresql server
$ make docker-down
```

#### DB Migration

```bash
# if you want to init your schema:
$ make migrate-up

# if you want to remove your schema:
$ make migrate-down
```

### Start HTTP server

```bash
$ go run main.go server
```

### Example API requests

```bash
# count
curl "http://127.0.0.1:8080/api/v1/it-test/count"

# list
curl "http://127.0.0.1:8080/api/v1/it-test/user/list?pageIndex=0&limit=2&orderBy=id&order=asc"

# create fail
cat << EOT > /tmp/create_user.json
{
  "aszf": true,
  "email": "test@test.com",
  "firstName": "Test",
  "lastName": "Test",
  "mobile": "06201234567",
  "password": "kisnyuszi",
  "passwordCheck": "a",
  "userName": "Test"
}
EOT
curl -X POST -d @/tmp/create_user.json "http://127.0.0.1:8080/api/v1/it-test/user"

# create ok
cat << EOT > /tmp/create_user.json
{
  "aszf": true,
  "email": "test@test.com",
  "firstName": "Test",
  "lastName": "Test",
  "mobile": "06201234567",
  "password": "kisnyuszi",
  "passwordCheck": "kisnyuszi",
  "userName": "Test"
}
EOT
curl -X POST -d @/tmp/create_user.json "http://127.0.0.1:8080/api/v1/it-test/user"

# update ok
cat << EOT > /tmp/update_user.json
{
  "firstName": "Test2",
  "lastName": "Test2",
  "mobile": "06301234567",
  "password": "kisnyuszi2",
  "passwordCheck": "kisnyuszi2",
  "userName": "Test2"
}
EOT
curl -X PUT -d @/tmp/update_user.json "http://127.0.0.1:8080/api/v1/it-test/user/cb3f0e3b-8de2-49b1-bfea-784c798a4b8e"
```
