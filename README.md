
# microapp
**How to run**

Execute migration for the first time

    go run app/migration/sqlite_initial_db.go 

Start the program

    go run main.go

**API Endpoints**

POST http://127.0.0.1:3000/register
POST http://127.0.0.1:3000/login
POST http://127.0.0.1:3000/login-history
GET http://127.0.0.1:3000/login-history
Sample postman requests can be imported from 
*login history app api.postman_collection.json*