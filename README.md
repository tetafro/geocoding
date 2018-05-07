# Geocoding

Search addresses and places.

## Database

Run PostgreSQL+PostGIS server
```sh
docker run -d \
    --name postgis \
    --publish 127.0.0.1:5432:5432 \
    --env 'POSTGRES_USER=postgres' \
    --env 'POSTGRES_PASSWORD=postgres' \
    --env 'POSTGRES_DB=geoplaces' \
    tetafro/postgis:10-2.4
```

Install `hstore` extension
```sh
docker exec -it postgis psql -U postgres geoplaces \
    -c 'CREATE EXTENSION hstore'
```

## Prepare data

Get map in PBF format
```sh
curl --output data/south-fed-district-latest.osm.pbf \
    https://download.geofabrik.de/russia/south-fed-district-latest.osm.pbf
```

Use [imposm](https://github.com/omniscale/imposm3/releases) importer for
[OpenStreetMap](https://www.openstreetmap.org/) data
```sh
imposm import \
    -mapping ./data/mapping.yml \
    -read ./data/south-fed-district-latest.osm.pbf \
    -connection postgis://postgres:postgres@localhost:5432/geoplaces \
    -overwritecache \
    -write
```

Init DB table and copy raw data to it
```sh
docker exec -i postgis psql -U postgres -h localhost geoplaces < ./data/init.sql
docker exec -i postgis psql -U postgres -h localhost geoplaces < ./data/copy.sql
docker exec -i postgis psql -U postgres -h localhost geoplaces < ./data/clean.sql
```

## Run

Copy and populate config
```sh
cp config.env.example config.env
```

Build and run the app
```sh
make dep build
source config.env
./bin/geocoding
```

## Try it

Open [http://localhost:8080](http://localhost:8080) in your browser or use `curl`
```sh
curl --request GET --url 'http://localhost:8080/api/v1/places?name=lenina%20street'
```
