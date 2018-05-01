# Geocoding

Search addresses and places.

## Prepare data

Use [OSMNames](https://github.com/OSMNames/OSMNames) project to extract data
from [OpenStreetMaps](https://www.openstreetmap.org/)
```sh
docker-compose run --rm osmnames
```

And take result files
```sh
cp OSMNames/data/export/kaliningrad-latest_geonames.tsv place.tsv
cp OSMNames/data/export/kaliningrad-latest_housenumbers.tsv address.tsv
```

## Database

Run Postgres
```sh
docker run -d \
    --name pg \
    --publish 127.0.0.1:5432:5432 \
    --volume $(pwd):/data \
    --env POSTGRES_USER=pguser \
    --env POSTGRES_PASSWORD=123 \
    --env POSTGRES_DB=geoplaces \
    postgres:10.3
```

Create tables structure and import data from tsv-files
```
docker exec -it pg psql -U pguser -h localhost geoplaces
\i /data/scripts/init.sql
\i /data/scripts/copy.sql
```

## Run

Copy and populate config
```sh
cp config.env.example config.env
```

Build and run the app
```sh
make build
source config.env
./bin/geocoding
```

## Try it

```sh
curl --request GET --url 'http://localhost:8080/api/places?name=lenina%20street'
```
