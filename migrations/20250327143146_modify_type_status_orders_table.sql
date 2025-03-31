-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
MODIFY COLUMN order_status INT(5) DEFAULT 0
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
MODIFY COLUMN order_status VARCHAR(255);
-- +goose StatementEnd
