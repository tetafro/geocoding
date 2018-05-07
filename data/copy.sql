-- Import point objects
INSERT INTO place (
    osm_id, osm_type, name,
    country, city, street, housenumber,
    coordinate
)
SELECT osm_id, all_tags -> 'type',
    COALESCE(NULLIF(name_ru, ''), name),
    country, city, street, housenumber,
    POINT(
        ST_X(ST_Transform(geometry, 4326)),
        ST_Y(ST_Transform(geometry, 4326))
    )
FROM import.osm_point;

-- Import linestring objects
INSERT INTO place (
    osm_id, osm_type, name,
    country, city, street, housenumber,
    coordinate
)
SELECT osm_id, all_tags -> 'type',
    COALESCE(NULLIF(name_ru, ''), name),
    country, city, street, housenumber,
    POINT(
        ST_X(ST_Transform(geometry, 4326)),
        ST_Y(ST_Transform(geometry, 4326))
    )
FROM (
    SELECT osm_id, all_tags, name, name_ru,
        country, city, street, housenumber,
        ST_Centroid(geometry) AS geometry
    FROM import.osm_linestring
) AS t;

-- Import polygon objects
INSERT INTO place (
    osm_id, osm_type, name,
    country, city, street, housenumber,
    coordinate
)
SELECT osm_id, all_tags -> 'type',
    COALESCE(NULLIF(name_ru, ''), name),
    country, city, street, housenumber,
    POINT(
        ST_X(ST_Transform(geometry, 4326)),
        ST_Y(ST_Transform(geometry, 4326))
    )
FROM (
    SELECT osm_id, all_tags, name, name_ru,
        country, city, street, housenumber,
        ST_Centroid(geometry) AS geometry
    FROM import.osm_polygon
) AS t;

-- Assemble fullname
UPDATE place
SET fullname = CONCAT_WS(
    ', ',
    NULLIF(city, ''),
    NULLIF(street, ''),
    NULLIF(housenumber, ''),
    NULLIF(name, '')
);

-- Prepare data for full-text search
UPDATE place SET tsv = to_tsvector(coalesce(fullname, ''));
