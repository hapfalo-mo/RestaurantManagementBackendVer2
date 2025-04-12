-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders DROP FOREIGN KEY orders_ibfk_1;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE booking DROP FOREIGN KEY booking_ibfk_1;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS user;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
