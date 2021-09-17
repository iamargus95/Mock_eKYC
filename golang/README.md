## Implementation as of 17.09.2021

### Initial state of the DB.

- i.e All data within the database.
![localhost:8080/albums](./images/initialStateofDB.png)

1. Run the program using `go run main.go`
2. Use PostMan for the following operations.


## 1. Get all data within the database.

- [GET] localhost:8080/albums
![postman_get_1](./images/BasicGetRequest.png)

## 2. Get a particular row of data using it's ID.

- [GET] localhost:8080/album/3
![postman_get_2](./images/GetSingleRow.png)

## 3. Edit the Price of the album using a PUT request.

- [PUT] localhost:8080/album/3
![postman_put_1](./images/EditPriceUsingID.png)

Here is the database after a successful PUT operation.
![DB](./images/databaseAfterFirstAlteration.png)

## 4. Adding a new row to the database using POST.

- [POST] localhost:8080/album
![postman_post](./images/newEntry.png)

- Here is the Database after the POST operation
![db_post](./images/dbAfterPost.png)

## 5. DELETE a row using it's ID.

- [DELETE] localhost:8080/album/2
![postman_delete](./images/postmanDelete.png)

- Database post Delete Operation
![db_delete](./images/dbDelete.png)