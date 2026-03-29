# dogs-api-Project

## Description

This project is a simple CRUD-based API for managing dog data. It consists of:

* A **[generator service](generator)** that loads initial data from a JSON file ([_dogs.json_](generator/dogs.json)) into a MySQL database.
* An **[API service](api)** that exposes endpoints to create, read, update, and delete dog records.
* A **[local environment setup](local-env)** using Docker Compose to run MySQL.

The system is designed to demonstrate a typical backend service structure using Go, Fiber, and MySQL.
## Project Structure

```
├── api/                # API service
├── common/             # It is a shared folder that contains module which can be used by any service(api/generator).
├── generator/          # Data loader (JSON → DB)
├── local-env/          # Docker setup for MySQL
├── .gitignore
└── README.md
```