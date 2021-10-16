CREATE DATABASE quizapp_dev;

CREATE TABLE IF NOT EXISTS quiz
(
    id serial PRIMARY KEY,
    description varchar NOT NULL UNIQUE,
    answer varchar NOT NULL
);
