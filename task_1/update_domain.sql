CREATE OR REPLACE FUNCTION update_domain()
    RETURNS void AS
$$
DECLARE
spot RECORD;
BEGIN
    RAISE NOTICE 'Getting row with domains';

FOR spot IN SELECT * FROM "MY_TABLE" WHERE website IS NOT NULL LOOP
UPDATE "MY_TABLE"
SET website = substring(spot.website, '(?![www.])(?:[a-z0-9-]{0,61}\.)+[a-zA-Z0-9-]{0,61}')
WHERE id = spot.id;
END LOOP;

    RETURN;
END;
$$ LANGUAGE plpgsql;