
# microapp
**How to run**

Execute migration for the first time

    go run app/migration/sqlite_initial_db.go 

Start the program

    go run main.go

**API Endpoints**

POST http://127.0.0.1:3000/register  <br>
POST http://127.0.0.1:3000/login  <br>
POST http://127.0.0.1:3000/login-history  <br>
GET http://127.0.0.1:3000/login-history  <br>
Sample postman requests can be imported from  <br>
*login history app api.postman_collection.json*
