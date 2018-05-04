# Geocoding

Search addresses and places.

## Prepare data

Use [OSMNames](https://github.com/OSMNames/OSMNames) project to extract data
from [OpenStreetMaps](https://www.openstreetmap.org/). Fix name field in
mapping if you want to use another language
```sh
sed -i 's/name:en/name:ru/g' data/import/mapping.yaml
```

Generate data files
```sh
docker-compose run --rm osmnames
```

And take results
```sh
cp OSMNames/data/export/kaliningrad-latest_geonames.tsv place.tsv
cp OSMNames/data/export/kaliningrad-latest_housenumbers.tsv address.tsv
```

## Database

Run Postgres
```sh
docker run -d \
    --name postgres-geocoding \
    --publish 127.0.0.1:5432:5432 \
    --volume $(pwd):/data \
    --env POSTGRES_USER=pguser \
    --env POSTGRES_PASSWORD=123 \
    --env POSTGRES_DB=geoplaces \
    postgres:10.3
```

Create tables structure and import data from tsv-files
```
docker exec -it postgres-geocoding psql -U pguser -h localhost geoplaces
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

Open [http://localhost:8080](http://localhost:8080) in your browser or use `curl`
```sh
curl --request GET --url 'http://localhost:8080/api/v1/places?name=lenina%20street'
```
