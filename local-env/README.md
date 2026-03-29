# local-env

## What This Does
Spins up a local MySQL database using Docker

## Steps

### 1. Start MySQL using Docker

```bash
cd local-env

docker compose build
docker compose up -d
```


### 2. Verify MySQL is Running

```bash
docker exec -it my-local-mysql /bin/bash
mysql -u root -p
```

Enter password:

```text
hello987
```

Then run:

```sql
show databases;
```

If this fails, your environment is not set up correctly. Fix that before moving on.
