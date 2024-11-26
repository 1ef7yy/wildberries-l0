-- +goose Up
-- +goose StatementBegin
ALTER TABLE data ADD CONSTRAINT order_uid_primary PRIMARY KEY (OrderUid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE data DROP CONSTRAINT order_uid_primary;
-- +goose StatementEnd
