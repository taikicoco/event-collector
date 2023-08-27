-- +goose Up
CREATE TABLE log_details (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    log_name_id BIGINT UNSIGNED,
    version BIGINT UNSIGNED,
    created_at TIMESTAMP null,
    updated_at TIMESTAMP null,
    PRIMARY KEY(id),
    FOREIGN KEY (log_name_id) REFERENCES log_names(id)
);
