#! /bin/sh

# reset test db, build sqlboiler models, etc.

DB_NAME="gelatin_test"
POSTGRESQL_URL="postgres://otto:letmeinnow@localhost:5432/${DB_NAME}?sslmode=disable"


echo "Resetting the test database 👯."
if [ "$( psql -d "postgres" -tAc "SELECT 1 FROM pg_database WHERE datname='$DB_NAME'" )" = '1' ]
then
    dropdb ${DB_NAME}
fi
    
createdb ${DB_NAME}
psql -d "postgres" -c "GRANT ALL ON DATABASE $DB_NAME TO otto;"

echo "Running 🏃migrations"
migrate -database ${POSTGRESQL_URL} -path ../db/migrations up

echo "Generating sqlboiler models 🧐"
sqlboiler --config="../db/sqlboiler-dev.toml" psql


echo "Running tests 🏁"
# Use Richgo if available
if command -v richgo &> /dev/null
then
  richgo test -v ./...
else
  go test -v ./...
fi
 