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
    if [ "$( psql -d "postgres" -tAc "SELECT 1 FROM pg_database WHERE datname='$DB_NAME'" )" = '1' ]
    then
        dropdb ${DB_NAME}
    fi
    
    createdb ${DB_NAME}
    psql -d "postgres" -c "GRANT ALL ON DATABASE $DB_NAME TO otto;"
fi

echo "Running 🏃migrations"
migrate -database ${POSTGRESQL_URL} -path ../db/migrations up

echo "Generating sqlboiler models 🧐"
sqlboiler --config="../db/sqlboiler-dev.toml" psql
