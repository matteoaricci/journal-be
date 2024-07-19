-- +goose Up
CREATE TABLE journals (
    id int NOT NULL,
    title text,
    body text,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE journals;
