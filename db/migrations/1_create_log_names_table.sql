-- +goose Up
CREATE TABLE log_names (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    name VARCHAR(255),
    PRIMARY KEY(id)
);
