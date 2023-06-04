-- +goose Up
-- +goose StatementBegin

CREATE TABLE logs (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    log_detail_id BIGINT UNSIGNED,
    access INTEGER,
    conversion INTEGER,
    created_at TIMESTAMP null,
    updated_at TIMESTAMP null,
    PRIMARY KEY(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'drop table logs';
-- +goose StatementEnd
