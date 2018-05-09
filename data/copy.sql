-- Delete unusable data
DELETE FROM import.osm_point
WHERE ST_IsEmpty(geometry) OR (
    name = '' AND
    name_ru = '' AND
    city = '' AND
    street = '' AND
    housenumber = ''
);
DELETE FROM import.osm_linestring
WHERE ST_IsEmpty(geometry) OR (
    name = '' AND
    name_ru = '' AND
    city = '' AND
    street = '' AND
    housenumber = ''
);
DELETE FROM import.osm_polygon
WHERE ST_IsEmpty(geometry) OR (
    name = '' AND
    name_ru = '' AND
    city = '' AND
    street = '' AND
    housenumber = ''
);

-- Import point objects
INSERT INTO place (
    osm_id, country, city, street, housenumber, name, type, coordinate
)
SELECT osm_id,
    country, city, street, housenumber,
    COALESCE(NULLIF(name_ru, ''), name),
    POINT(
        ST_X(ST_Transform(geometry, 4326)),
        ST_Y(ST_Transform(geometry, 4326))
    )
FROM import.osm_point;

-- Import linestring objects
INSERT INTO place (
    osm_id, country, city, street, housenumber, name, type, coordinate
)
SELECT osm_id,
    country, city, street, housenumber,
    COALESCE(NULLIF(name_ru, ''), name),
    POINT(
        ST_X(ST_Transform(geometry, 4326)),
        ST_Y(ST_Transform(geometry, 4326))
    )
FROM (
    SELECT osm_id,
        country, city, street, housenumber,
        name, name_ru,
        all_tags,
        ST_Centroid(geometry) AS geometry
    FROM import.osm_linestring
) AS t;

-- Import polygon objects
INSERT INTO place (
    osm_id, country, city, street, housenumber, name, type, coordinate
)
SELECT osm_id,
    country, city, street, housenumber,
    COALESCE(NULLIF(name_ru, ''), name),
    POINT(
        ST_X(ST_Transform(geometry, 4326)),
        ST_Y(ST_Transform(geometry, 4326))
    )
FROM (
    SELECT osm_id,
        country, city, street, housenumber,
        name, name_ru,
        all_tags,
        ST_Centroid(geometry) AS geometry
    FROM import.osm_polygon
) AS t;
