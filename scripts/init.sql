CREATE TABLE place
(
    id                SERIAL PRIMARY KEY,
    name              VARCHAR,
    alternative_names VARCHAR,
    osm_type          VARCHAR,
    osm_id            BIGINT,
    class             VARCHAR,
    type              VARCHAR,
    lon               DOUBLE PRECISION,
    lat               DOUBLE PRECISION,
    place_rank        INT,
    importance        DOUBLE PRECISION,
    street            VARCHAR,
    city              VARCHAR,
    county            VARCHAR,
    state             VARCHAR,
    country           VARCHAR,
    country_code      CHAR(2),
    display_name      VARCHAR,
    west              DOUBLE PRECISION,
    south             DOUBLE PRECISION,
    east              DOUBLE PRECISION,
    north             DOUBLE PRECISION,
    wikidata          VARCHAR,
    wikipedia         VARCHAR,
    housenumbers      VARCHAR
    tsv               tsvector;
);
CREATE INDEX tsv_idx ON place USING gin(tsv);

CREATE TABLE address
(
    id SERIAL PRIMARY KEY,
    osm_id      BIGINT,
    street_id   BIGINT,
    street      VARCHAR,
    housenumber VARCHAR,
    lon         DOUBLE PRECISION,
    lat         DOUBLE PRECISION
    tsv         tsvector;
);
CREATE INDEX tsv_idx ON address USING gin(tsv);
