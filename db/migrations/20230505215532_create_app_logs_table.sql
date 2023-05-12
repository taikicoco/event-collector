-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    user_id VARCHAR(255),
    password VARCHAR(255),
    nickname VARCHAR(255) null,
    comment VARCHAR(255) null,
    created_at TIMESTAMP,
    updated_at TIMESTAMP null,
    deleted_at TIMESTAMP null,
    PRIMARY KEY(user_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
