-- +goose Up
-- +goose StatementBegin
ALTER TABLE history
DROP COLUMN updated_at;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE history
ADD COLUMN updated_at TIMESTAMP NOT NULL;
-- +goose StatementEnd
