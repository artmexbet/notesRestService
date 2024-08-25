# Notes Rest Service
## Installation
1. Clone the repository
2. Run `docker build -t "notes_build" -f .\deploy\Dockerfile .` in console
3. Run `docker-compose up -d`
## Usage
You can import Notes.postman_collection.json to Postman to test the service.

There are 4 endpoints:
1. GET /notes - get all notes
2. POST /notes - create a note
3. POST /register - register a user
4. POST /login - login a user

## Technologies
- Go
- Chi
- pgx
- Docker
- Postman
- PostgreSQL
- JWT