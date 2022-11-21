CREATE OR REPLACE FUNCTION update_domain()
    RETURNS void AS
$$
DECLARE
    spot RECORD;
    domains text[];
    pattern text;
BEGIN
    pattern := '(:?www.)?((?:[a-z0-9-]+\.)+(?:[a-z]+))(?:\/)?';

    FOR spot IN SELECT * FROM "MY_TABLE" WHERE website ~ pattern LOOP
        SELECT ARRAY_AGG(a)
        INTO domains
        FROM (
            SELECT DISTINCT unnest(
                regexp_matches(spot.website, pattern, 'g')
            ) as a
        ) as a2;

        UPDATE "MY_TABLE"
        SET website = array_to_string(ARRAY_REMOVE(domains, 'www.'), ',')
        WHERE id = spot.id;
    END LOOP;

    RETURN;
END;
$$ LANGUAGE plpgsql;