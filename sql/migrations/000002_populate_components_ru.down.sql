DELETE FROM panel_components
WHERE
    language_id = (
        SELECT
            id
        FROM
            languages
        WHERE
            name = 'Русский'
    );

DELETE FROM panel_measurement_units
WHERE
    language_id = (
        SELECT
            id
        FROM
            languages
        WHERE
            name = 'Русский'
    );

DELETE FROM languages
WHERE
    name = 'Русский';
