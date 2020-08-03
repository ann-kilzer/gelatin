BEGIN;

CREATE TABLE IF NOT EXISTS locations(
   id serial PRIMARY KEY,
   location_name VARCHAR (100) UNIQUE NOT NULL,
   softness INTEGER NOT NULL
);

COMMIT;