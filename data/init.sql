CREATE TABLE place
(
    id                SERIAL PRIMARY KEY,
    osm_id            BIGINT,
    osm_type          VARCHAR,
    name              VARCHAR,
    coordinate        POINT,
    tsv               tsvector
);
CREATE INDEX tsv_idx ON place USING gin(tsv);
