INSERT INTO place (osm_id, osm_type, name, coordinate)
SELECT osm_id, all_tags -> 'type',
    COALESCE(NULLIF(name_ru, ''), name),
    POINT(ST_X(geometry), ST_Y(geometry))
FROM import.osm_point;

INSERT INTO place (osm_id, osm_type, name, coordinate)
SELECT osm_id, all_tags -> 'type',
    COALESCE(NULLIF(name_ru, ''), name),
    POINT(ST_X(geometry), ST_Y(geometry))
FROM (
    SELECT osm_id, all_tags, name, name_ru,
        ST_Centroid(geometry) AS geometry
    FROM import.osm_linestring
) AS t;

INSERT INTO place (osm_id, osm_type, name, coordinate)
SELECT osm_id, all_tags -> 'type',
    COALESCE(NULLIF(name_ru, ''), name),
    POINT(ST_X(geometry), ST_Y(geometry))
FROM (
    SELECT osm_id, all_tags, name, name_ru,
        ST_Centroid(geometry) AS geometry
    FROM import.osm_polygon
) AS t;

UPDATE place SET tsv = to_tsvector(coalesce(name, ''));
