# go-dynamodb-crud-api
This repository aims to look at a different structure for buildng CRUD APIs with Go, Chi, and DynamoDB.

## ABOUT THE PROJECT

### Structure
This API uses Routes to talk to the Handlers, who then talk to the Controllers, who then talk to the repository/adapter on a database level.
There is also a health folder to check on the overall status/health of the API and database.

### "Entities"
Entities consist of two parts—a Base and a Product. The Base part of the entity holds generic fields such as ID, CreatedAt, UpdatedAt, etc.
The Product part holds any other fields unique to the item and nests the Base part inside.

### DynamoDB & Data
This API is interacted with via JSON data. This means the data passed in will be unmarshalled into Structs so Go will understand it.
From there, it will be converted into DynamoDB attributes (and vice versa).

## IMPORTANT NOTES
- This project uses older packages since it was put on the back burner half way into development back in 2022.
- This project was finished recently after wrapping up other projects in Python and JavaScript.
- This project may be updated in the future to fix any leftover bugs, add comments, or update packages.
