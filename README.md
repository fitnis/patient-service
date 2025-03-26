# Patient Service

This microservice handles patient registration and management for the Fitnis healthcare system.

## Features

- Patient registration
- Patient data management (CRUD operations)
- Integration with Keycloak for authentication and authorization
- Role-based access control

## API Endpoints

| Method | Endpoint             | Description            | Required Role          |
| ------ | -------------------- | ---------------------- | ---------------------- |
| GET    | /health              | Health check           | None                   |
| POST   | /api/v1/patients     | Register a new patient | Doctor                 |
| GET    | /api/v1/patients     | Get all patients       | Doctor                 |
| GET    | /api/v1/patients/:id | Get a specific patient | Any authenticated user |
| PUT    | /api/v1/patients/:id | Update a patient       | Doctor                 |
| DELETE | /api/v1/patients/:id | Delete a patient       | Doctor                 |

## Setup and Running

### Environment Variables
