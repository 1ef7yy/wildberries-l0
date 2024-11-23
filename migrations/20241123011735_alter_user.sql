-- +goose Up
-- +goose StatementBegin
ALTER ROLE "admin" WITH SUPERUSER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER ROLE "admin" WITH NOSUPERUSER;
-- +goose StatementEnd
