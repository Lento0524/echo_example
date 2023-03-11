CREATE TABLE food (
    id SERIAL NOT NULL PRIMARY KEY,
    name varchar NOT NULL,
    unit varchar NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);
