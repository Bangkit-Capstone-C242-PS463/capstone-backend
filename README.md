# Google Bangkit 2024 Capstone API Service
InSight API service

## InSight API Documentation
### Base URL: `<BASE_URL>/api`
| Method | Endpoint                 | Summary                                | Description                                                             | Request Body                       | Response                     |
|--------|--------------------------|----------------------------------------|-------------------------------------------------------------------------|------------------------------------|------------------------------|
| POST   | `/v1/auth/login`         | User login                             | Authenticates a user and returns data upon successful login.            | `dto.LoginRequest`                | `dto.LoginResponse`          |
| POST   | `/v1/auth/signup`        | Sign up a new user                     | Registers a new user in the system.                                     | `dto.SignUpRequest`               | `dto.SuccessResponse`        |
| GET    | `/v1/health`             | Get server health status               | Returns the current health status of the server.                        | None                               | `dto.ServerHealthResponse`   |
| POST   | `/v1/predict`            | Predict disease based on user story    | Predicts the disease based on the provided user story.                  | `dto.PredictRequest`              | `dto.PredictResponse`        |
| POST   | `/v1/predict_manual`     | Predict disease based on symptoms      | Predicts the disease based on the provided symptoms.                    | `dto.PredictManualRequest`         | `dto.PredictResponse`        |
| GET    | `/v1/user/history`       | Get user diagnosis history             | Retrieves the history of user diagnoses based on their submitted symptoms. | None                               | `dto.GetAllHistoryResponse`  |
| DELETE | `/v1/user/history`       | Delete a diagnosis history record      | Deletes a specific diagnosis history record by its ID.                  | Query Param: `history_id`          | `dto.SuccessResponse`        |

### Notes:
- **Request Bodies**: All request bodies follow the respective data transfer objects (`dto`) as defined in the Swagger schema (see `docs/swagger.json`).
- **Response**: Responses adhere to standard objects (`dto`) for success, errors, or other relevant data.
- **Error Responses**:
  - `400`: Bad Request - Invalid input.
  - `401`: Unauthorized - Invalid credentials.
  - `404`: Not Found - User or resource not found.
  - `500`: Internal Server Error - General server errors.

## Project Setup
Follow these steps to set up and run the application:

### Setup Environment
#### 1. Copy the `.env` structure to create your development environment file:
```bash
cp .env-sample .env.dev
```
Update the values in `.env.dev` as needed.

#### 2. Run the Application
Use the `Makefile` to start the application in development mode:
```bash
make start-dev
```
The server will start on port `5000`

### Database Migrations
#### Manage database migrations with the following commands:

* Migrate Up

To apply all pending migrations and update the database schema:
```bash
make migrate-up
```

* Migrate Down

To roll back the most recent migration:
```bash
make migrate-down
```

### API Documentation
#### Generating API Documentation
If you add a new endpoint or update Swagger annotations, you can regenerate the API documentation using the following command:
```bash
make generate-docs
```

#### Accessing API Documentation
Once the application is running, the API documentation is available at:
```bash
/api/v1/swagger/index.html
```
Open this URL in your browser to view the Swagger UI and explore the available endpoints.

### Credits
This project is implemented by Team C242-PS463.
