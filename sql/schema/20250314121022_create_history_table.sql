-- +goose Up
-- +goose StatementBegin
CREATE TABLE history(
    id UUID PRIMARY KEY,
    point INTEGER NOT NULL,
    reason TEXT NOT NULL,
    date INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE history;
-- +goose StatementEnd
