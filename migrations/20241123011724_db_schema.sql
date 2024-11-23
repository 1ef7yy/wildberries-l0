-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS data (
    OrderUid VARCHAR(100),
    Data JSONB
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE data CASCADE;
-- +goose StatementEnd


