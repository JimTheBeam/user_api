# user_api is a small api that provide work with users

Server can use two different types of storages:

* postgres database
* json object file

You can choose one of them in the config file.

## Methods

Parameter content type: application/json

1. CREATE a new user:<br>
Method: POST<br>
Path: /v1/user/<br>
Request json: {"name": "string"}<br>
Successful response json: {"id": integer, "name": "string", "created_at": "string"}<br>

2. GET a USER with ID.<br>"
Method: GET<br>
Path: /v1/user/"id"<br>
Successful response json: {"id": integer, "name": "string", "created_at": "string"}<br>

3. GET ALL users.<br>
Method: GET<br>
Path: /v1/user/users<br>
Successful response json: [{"id": integer, "name": "string", "created_at": "string"}, {},...]<br>

4. DELETE a user with id<br>
Method: DELETE<br>
Path: /v1/user/"id"<br>
Successful response json: { "code": 200, "name": "OK", "message": "OK"}<br>

5. UPDATE a user with id.<br>
Method: PUT<br>
Path: /v1/user/"id"<br>
Request json: {"name": "string", id: integer}<br>
Successful response json: {"id": integer, "name": "string", "created_at": "string"}<br>

**Error response json for all methods:**<br>
json: { "code": 400, "name": "string", "message": "string"}<br>

**Json fields:**<br>
id - user id; <br>
name - user name; <br>
created_at - created time, format: "yyyy-mm-ddThh:mm:ss";<br>

## To run the server

1. Create file config/env.sh

2. Add settings in the file:

> server:<br>
> export HTTP_ADDR=localhost:8080<br>
> export READ_TIMEOUT="30s"<br>
> export WRITE_TIMEOUT="30s"<br>
>
> path to log file:<br>
> export LOG_FILE_PATH="./log/log.log"<br>
>
> storage settings. Can be only "postgres" or "jsonObj"<br>
> export STORAGE="jsonObj"<br>
>
> jsonObj settings:<br>
> export JSON_FILE_PATH="users.json"<br>
>
> database settings:<br>
> export DB_USERNAME="postgres"<br>
> export DB_HOST="localhost"<br>
> export DB_PORT="database port"<br>
> export DB_NAME="user_api"<br>
> export DB_SSL_MODE="disable"<br>
> export DB_PASSWORD="database password"<br>

3. Run the program with command: ". ./rundev.sh"
