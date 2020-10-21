DROP TABLE IF EXISTS samples;

CREATE TABLE IF NOT EXISTS samples
(
   id SERIAL NOT NULL,
   text TEXT NOT NULL,
   PRIMARY KEY (id)
);

INSERT INTO samples(text) VALUES ('sample');
