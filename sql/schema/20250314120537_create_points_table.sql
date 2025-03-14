-- +goose Up
-- +goose StatementBegin
CREATE TABLE points(
    id UUID PRIMARY KEY,
    total INTEGER,
    goal INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE points;
-- +goose StatementEnd
