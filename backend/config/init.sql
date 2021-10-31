CREATE TABLE IF NOT EXISTS quiz
(
    id serial PRIMARY KEY,
    description varchar NOT NULL UNIQUE,
    answer varchar NOT NULL,
    user_id integer NOT NULL
);

CREATE TABLE IF NOT EXISTS users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);