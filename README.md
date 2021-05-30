# user_api is a small api that works with users.

## Methods:

POST host/v1/user/ - create new user

GET host/v1/user/<id> - get user 

GET host/v1/user/users - get all users

GET host/v1/user/delete/<id> - delete user

PUT host/v1/user/<id> - update user


## To run the server:
1. Create file config/env.sh

 2. Add settings in the file:

 settings:

    host:
    export HTTP_ADDR=localhost:8080

    path to log file:
    export LOG_FILE_PATH="./log/log.log"

    database settings:
    export DB_USERNAME=<postgres>
    export DB_HOST=<localhost>
    export DB_PORT=<database port>
    export DB_NAME="user_api"
    export DB_SSL_MODE="disable"
    export DB_PASSWORD=<password>


3. run the program with command: ". ./rundev.sh"