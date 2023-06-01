-- +goose Up
-- +goose StatementBegin

CREATE TABLE logs (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    log INTEGER,
    created_at TIMESTAMP null,
    updated_at TIMESTAMP null,
    PRIMARY KEY(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'drop SQL query';
-- +goose StatementEnd
