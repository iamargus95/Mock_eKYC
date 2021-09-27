# Pre-Requisites

1. Create a database in PostgreSQL with the name "mock_ekyc" with the username "postgres".
2. Keep the password as "abc123"
3. Create a .env file with the following contents
    -   ```export DIALECT="postgres"
        export HOST="localhost"
        export DBPORT="5432"
        export DBUSER="postgres"
        export DBNAME="mock_ekyc"
        export PASSWORD="abc123"
        ```

## Build the program

1. Open a terminal inside `/golang/eKYC-service-gin`
2. Use the command `go build .`
3. then do `./eKYC-service-gin`

## Use Postman to perform the post operation

1. Perform the POST operation at `http://localhost:8080/api/v1/signup` with a 
- request body 

    ```JSON
    {
        "name":"some-name",
        "email":"some-email", // Only Valid emails allowed.
        "plan":"some-plan" //One of basic, advanced, enterprise.
    }
    ```