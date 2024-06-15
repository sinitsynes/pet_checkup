CREATE TABLE panel_units (
	id integer PRIMARY KEY generated by default as identity,
	name text NOT NULL,
	language_id integer references languages(id)
);

CREATE TABLE panel_components (
	id integer PRIMARY KEY generated by default as identity,
	name text NOT NULL,
	unit_id integer references panel_units(id),
	average_amount numrange,
	comment text,
	language_id integer references languages(id)
);

CREATE TABLE panel_measurements (
	id integer PRIMARY KEY generated by default as identity,
	component_id integer references panel_components(id),
	date_measured date NOT NULL,
	pet_id integer references pets(id),
	amount real NOT NULL);

