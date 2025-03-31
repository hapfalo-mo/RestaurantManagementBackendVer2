-- +goose Up
-- +goose StatementBegin
ALTER TABLE `orders`
ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL,
ADD COLUMN note VARCHAR(250),
ADD COLUMN feedback INT(5) 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `orders`
DROP COLUMN updated_at,
DROP COLUMN deleted_at,
DROP COLUMN note,
DROP COLUMN feedback;
-- +goose StatementEnd
