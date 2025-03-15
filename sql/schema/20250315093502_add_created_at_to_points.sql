-- +goose Up
-- +goose StatementBegin
ALTER TABLE points
ADD COLUMN created_at TIMESTAMP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE points
DROP COLUMN created_at;
-- +goose StatementEnd
