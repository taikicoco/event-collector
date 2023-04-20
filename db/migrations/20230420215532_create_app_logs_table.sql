-- +goose Up
-- +goose StatementBegin

CREATE TABLE app_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    log_key VARCHAR(300),
    log_value VARCHAR(3000),
    created_at TIMESTAMP,
    updated_at TIMESTAMP null,
    deleted_at TIMESTAMP null,
    PRIMARY KEY(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
