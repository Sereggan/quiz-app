CREATE
DATABASE quizapp_dev;

CREATE TABLE quiz
(
    id          serial PRIMARY KEY,
    description varchar NOT NULL,
    answer      varchar NOT NULL
);

