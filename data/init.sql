CREATE TABLE place
(
    id           SERIAL PRIMARY KEY,
    osm_id       BIGINT,
    country      VARCHAR,
    city         VARCHAR,
    street       VARCHAR,
    housenumber  VARCHAR,
    name         VARCHAR,
    fullname     VARCHAR,
    coordinate   POINT,
    tsv_fullname tsvector
);
CREATE INDEX tsv_idx ON place USING gin(tsv_fullname);

CREATE OR REPLACE FUNCTION update_fullname()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.fullname = CONCAT_WS(
        ', ',
        NULLIF(NEW.city, ''),
        NULLIF(NEW.street, ''),
        NULLIF(NEW.housenumber, ''),
        NULLIF(NEW.name, '')
    );
    NEW.tsv_fullname = to_tsvector(COALESCE(NEW.fullname, ''));
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_place_fullname_update
    BEFORE INSERT OR UPDATE OF city, street, housenumber, name ON place
    FOR EACH ROW
    EXECUTE PROCEDURE update_fullname();
