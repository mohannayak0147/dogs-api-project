# api

## Description

The API service exposes REST endpoints for managing dog data.

It is built using:

* Go
* Fiber framework
* MySQL backend

## How It Works

1. Loads configuration
2. Establishes database connection
3. Registers routes
4. Starts HTTP server


## Environment Setup

* Navigate to the [local-env](../local-env) folder and follow as mentioned, to spin up a local MySQL database using Docker - ignore if already done. 

* Before running the API, create the environment file:

    ```bash
    cd api
    cp .env-dist .env
    ```

    ⚠️ `.env` is ignored by git and must not be committed.

    If this file is missing or wrong, the API will not start.


## Running the API

```bash
go run main.go
```

If the app crashes here, your DB connection is wrong. Don’t guess—check your config.


## API Endpoints

Base URL:

```
http://127.0.0.1:8081/api
```

### Health

* `GET /status`

### Dogs

* `GET /dogs` → Get all dogs
* `GET /dogs/:dog` → Get a specific dog
* `POST /dogs` → Create a dog
* `PUT /dogs/:dog` → Update a dog
* `DELETE /dogs/:dog` → Delete a dog


### Swagger Docs

```
http://127.0.0.1:8081/docs/index.html
```

## Notes

* API depends on the database being up and populated
* No data = useless API
* Always run generator first

