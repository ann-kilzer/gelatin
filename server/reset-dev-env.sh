#! /bin/sh

# reset db, build sqlboiler models, etc.

POSTGRESQL_URL='postgres://otto:letmeinnow@localhost:5432/gelatin?sslmode=disable'
DB_NAME="gelatin"

print_usage() {
    printf "Usage: This script rebuilds the backend"
    printf " "
    printf "$./reset-dev-env.sh [options]"
    printf " "
    printf "options:"
    printf "-h, --help                show brief help"
    printf "db                  drop and rebuild the database"
}

db_flag=false

while getopts 'abf:v' flag; do
  case "${flag}" in
    db) db_flag=true ;;
    *) print_usage
       exit 1 ;;
  esac
done


if [ $db_flag ]; then
    echo "Resetting the database 👯."
    dropdb ${DB_NAME}
    createdb ${DB_NAME}
    psql -U "otto" -d "postgres" -c "GRANT ALL ON DATABASE gelatin TO otto;"
fi

echo "Running 🏃migrations"
migrate -database ${POSTGRESQL_URL} -path ../db/migrations up

echo "Generating sqlboiler models 🧐"
sqlboiler --config="../db/sqlboiler-dev.toml" psql
