CREATE TABLE place
(
    id                SERIAL PRIMARY KEY,
    osm_id            BIGINT,
    osm_type          VARCHAR,
    coordinate        POINT,
    priority          SMALLINT DEFAULT 0 NOT NULL,
    name              VARCHAR,
    country           VARCHAR,
    city              VARCHAR,
    street            VARCHAR,
    housenumber       VARCHAR,
    fullname          VARCHAR,
    tsv               tsvector
);
CREATE INDEX tsv_idx ON place USING gin(tsv);
