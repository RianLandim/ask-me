-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS room (
    "id" uuid PRIMARY KEY DEFAULT  gen_random_uuid(),
    "theme" VARCHAR(255) NOT NULL
)

---- create above / drop below ----

DROP TABLE IF EXISTS room;
