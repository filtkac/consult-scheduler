# This project is a work in progress!

# Consult scheduler

This software was developed to help with scheduling consults and examinations in hospital settings.
It's a full-stack application with a React frontend, Go backend, and a PostgreSQL database.

## How to run

Currently, the application is not running in a production environment.
You can run the app locally by running the backend and frontend separately and setting up the database.

You need to create an `.env` file in the root directory and fill the necessary variables.
The list of variables to fill is in the `.env.example` file.

To set up the database, you can use the `docker-compose.yml` file.
When running the Go backend, you need to inject the environment variables from the `.env` file.

## Testing

Go backend has integration tests set up; all you need is a running docker agent, the integration tests set up the database using test containers.
There is also the `api.http` file which can be used by a JetBrains IDE to query the backend API.
The `test_data.sql` script contains some fake test data to fill in the database.
This test data is also used by the integration tests; after each test, the transaction is rolled back.
