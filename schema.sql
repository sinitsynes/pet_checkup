CREATE TABLE public.languages (
    id integer NOT NULL,
    name text NOT NULL
);


CREATE TABLE public.panels (
    id integer NOT NULL,
    unit_id integer,
    date_measured date,
    pet_id integer,
    amount integer
);


CREATE TABLE public.unit_names (
    id integer NOT NULL,
    name text
);


CREATE TABLE public.unit_translations (
    id integer NOT NULL,
    unit_id integer,
    language_id integer,
    comment text,
    name text
);


CREATE TABLE public.units (
    id integer NOT NULL,
    average_amount text,
    translation_id integer
);