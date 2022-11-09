CREATE TABLE IF NOT EXISTS accounts (
   id SERIAL PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL
);

INSERT INTO accounts (username, password) VALUES ('nico', '123456');

CREATE TABLE IF NOT EXISTS songs (
    id DOUBLE PRECISION NOT NULL,
    name TEXT,
    artist TEXT,
    duration DOUBLE PRECISION,
    album TEXT,
    artwork TEXT,
    price DOUBLE PRECISION,
    origin TEXT,
    CONSTRAINT songs_pkey PRIMARY KEY (id)
);