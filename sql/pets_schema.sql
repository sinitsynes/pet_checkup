CREATE TABLE pets (
    id integer PRIMARY KEY,
    name text NOT NULL,
    birth date,
    passport text,
    chip text
);


CREATE TABLE pet_weight (
    id integer PRIMARY KEY,
    weight real NOT NULL,
    date_measured date NOT NULL default CURRENT_DATE,
    pet_id integer NOT NULL
);