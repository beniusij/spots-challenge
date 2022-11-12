CREATE OR REPLACE FUNCTION get_all_with_domain()
    RETURNS SETOF "MY_TABLE" AS
$$
DECLARE
    spots "MY_TABLE";
BEGIN
    FOR spots IN
        SELECT * FROM "MY_TABLE" WHERE array_length(regexp_matches(website, ''), 1) > 0
    LOOP
        RETURN NEXT spots;
    END LOOP;
    RETURN;
END;
$$ LANGUAGE plpgsql;