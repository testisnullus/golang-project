CREATE TABLE current_weather (
                                 id SERIAL PRIMARY KEY,
                                 city VARCHAR(255) NOT NULL,
                                 description VARCHAR(255) NOT NULL,
                                 time TIMESTAMPTZ NOT NULL,
                                 speed DOUBLE PRECISION NOT NULL,
                                 temp DOUBLE PRECISION NOT NULL,
                                 humidity DOUBLE PRECISION NOT NULL
);