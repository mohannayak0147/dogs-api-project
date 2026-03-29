# generator

## Description

The generator service reads dog data from [**dogs.json**](dogs.json) and inserts it into the MySQL database.

This is a one-time or repeatable data seeding step.


## What It Does

* Opens `dogs.json`
* Parses dog records
* Inserts them into MySQL


## Environment Setup
* Navigate to the [local-env](../local-env) folder and follow as mentioned, to spin up a local MySQL database using Docker - ignore if already done.
* Before running the generator, create the environment file:

    ```bash
        cd generator
        cp .env-dist .env
    ```
    
     ⚠️ `.env` is ignored by git and must not be committed.
    
    Wrong or missing env = DB connection failure.


## Running the Generator

```bash
go run main.go
```

## Requirements

* MySQL container must be running
* Database must be accessible

If this step fails, don’t move forward. Your [API](../api) depends on this data.


## Common Mistakes

* Running generator before DB is up → fails
* Wrong DB credentials → fails
* Modifying JSON format → breaks parsing
