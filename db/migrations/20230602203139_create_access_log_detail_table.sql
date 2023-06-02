-- +goose Up
-- +goose StatementBegin

CREATE TABLE access_log_details (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    name VARCHAR(255),
    version INTEGER,
    page_url VARCHAR(255),
    created_at TIMESTAMP null,
    updated_at TIMESTAMP null,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'drop table access_log_details';
-- +goose StatementEnd
