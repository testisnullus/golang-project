CREATE TABLE users
(
    id         serial,
    first_name VARCHAR,
    last_name  VARCHAR,
    email      VARCHAR NOT NULL
);

CREATE TABLE hash_password
(
    user_id  INTEGER,
    salt     VARCHAR NOT NULL,
    password VARCHAR
);
