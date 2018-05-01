COPY place (
    name, alternative_names, osm_type, osm_id, class, type, lon, lat,
    place_rank, importance, street, city, county, state, country, country_code,
    display_name, west, south, east, north, wikidata, wikipedia, housenumbers
)
FROM '/data/place.tsv'
WITH CSV DELIMITER E'\t' HEADER;
UPDATE place SET tsv = to_tsvector(coalesce(display_name, ''));

COPY address (
    osm_id, street_id, street, housenumber, lon, lat
)
FROM '/data/address.tsv'
WITH CSV DELIMITER E'\t' HEADER;
UPDATE address SET tsv = to_tsvector(coalesce(street, ''));
