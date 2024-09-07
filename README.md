### HOW TO RUN THE PROJECT

1. Load the schema by running the `schema.sql` file and create database "auth"

2. Run the project:
   ```bash
   go run ./cmd/
   ```

### API ENDPOINTS

#### auth
1. **GET /auth/access**  
   Parameters:
    - `guid` (in query parameters)

2. **POST /auth/refresh**  
   Parameters:
    - `guid` (in request body)
    - `refresh token` (in request body)


### I decided to add my config to the git