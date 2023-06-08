-- +goose Up
-- +goose StatementBegin

CREATE TABLE web_log_data (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    web_log_id BIGINT UNSIGNED,
    access INTEGER,
    conversion INTEGER,
    created_at TIMESTAMP null,
    updated_at TIMESTAMP null,
    PRIMARY KEY(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'drop table web_log_data';
-- +goose StatementEnd
