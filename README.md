# Requirements

1. Signup a new api client to get access and secret keys.

2. Upload Images and Metadata to MINIO bucket.

3. Perform face matching between two image ids to get the face match score.

4. Perform Optical Character Recognition (OCR) on images.


# Pre-Requisites

1. Create a database in PostgreSQL with the name "mock_ekyc" with the username "postgres".
2. Keep the password as "abc123"
3. Create a .env file with the following contents

    -   ```
        export DIALECT="postgres"
        export HOST="localhost"
        export DBPORT="5432"
        export TESTDBPORT="2345"
        export DBUSER="postgres"
        export DBNAME="mock_ekyc"
        export PASSWORD="abc123"
        export MYSIGNINGKEY="jksvasvuyfglisnxcvjbvalifboashfosbfisbgfiuwgsfhdnvkbsljvsdi"
        export MINIOENDPOINT="localhost:9000"
        export MINIOACCESS="myaccesskey"
        export MINIOSECRET="mysecretkey"
        ```

## Build the program

1. Open a terminal inside `/golang/eKYC-service-gin`
2. Use the command `make build`
3. then do `make run`
