# gelatin

Golang skeleton app = Go + Skeleton => Gelatin

The stack includes:
- Vue.js frontend
- PostgreSQL DB
- Migrations
- SqlBoiler ORM
- Bash scripts to tie it all together!

## Getting started

Requirements:

- PostgreSQL 12.3
- Golang 1.14
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) installed as CLI tool
- Yarn v1.22.4


### Initial DB Setup

We create an "otto" user to get around OS X postgres weirdness. Replace with a strong password for production environments
```
psql
CREATE USER otto password 'letmeinnow';
CREATE DATABASE gelatin;
GRANT ALL ON DATABASE gelatin TO otto;
```

### Running the server

For the first time, you'll need to run the db migrations and build the ORM models. Add the `db` flag if you need a fresh database (destructive so use carefully)
```
cd server
./reset-dev-env.sh db
```

Now build and run the server with
```
./run.sh
```

### Running the frontend

```
cd web
yarn install
yarn run start
```

