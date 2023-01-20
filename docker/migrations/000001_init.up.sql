CREATE TABLE people
(
    id              SERIAL UNIQUE,
    name            VARCHAR(255) NOT NULL,
    phone_number       INTEGER,
    birth_year    INTEGER,
    nationality           VARCHAR(255)  NOT NULL
);