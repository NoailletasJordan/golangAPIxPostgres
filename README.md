# Golang API x Postgres

This repo helps kickstart projects by presenting an easily editable API with a [3-layered architecture](https://blog.jordannoailletas.com/en/published/three-layered-architecture), connected to a PostgreSQL database on Docker.

![showcase](https://github.com/user-attachments/assets/04564f61-9d8b-421d-9668-af3bdf9c18fc)

## Features

- Standard library `http` package
- Input validation
- Swagger documentation
- Unit tests
- Database interface on port 5433 (PgAdmin)

## Environment Variables

Create a `.env` file at the root with the following variables:

```env
DATABASE_FOLDER_NAME=
POSTGRES_PASSWORD=
POSTGRES_USER=
POSTGRES_DB=
POSTGRES_DB_TEST=
POSTGRES_EXTERNAL_PORT=
POSTGRES_EXTERNAL_HOST=
PGADMIN_DEFAULT_EMAIL=
```
