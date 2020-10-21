DROP TABLE IF EXISTS samples;

CREATE TABLE IF NOT EXISTS samples
(
    id SERIAL NOT NULL,
    text VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO samples(text) VALUES ('sample');
