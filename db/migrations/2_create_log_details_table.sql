-- +goose Up
CREATE TABLE log_details (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    log_name_id BIGINT UNSIGNED,
    version BIGINT UNSIGNED,
    PRIMARY KEY(id),
    FOREIGN KEY (log_name_id) REFERENCES log_names(id)
);
