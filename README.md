# goteam

# Migrating the development database

Run the following command to migrate to the latest schema:

Log into the gitlab docker registry:
`docker login registry.gitlab.com`

```
# for development
docker-compose -f docker-compose.yml -f docker-compose.db-migrate.yml run --rm trek-dev

# for testing
docker-compose -f docker-compose.yml -f docker-compose.db-migrate.yml run --rm trek-test
```

## Adding a new migration

```
docker-compose -f docker-compose.yml -f docker-compose.db-migrate.yml run --rm --no-deps trek-dev \
  ./dbtrek add --migrations /migrate/migrations.yml
```

The database by default will startup on localhost with default ports. Connection details are specified in the `docker-compose.yml`.
