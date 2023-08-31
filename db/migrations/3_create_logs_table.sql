-- +goose Up
CREATE TABLE logs (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    log_detail_id BIGINT UNSIGNED,
    created_at TIMESTAMP,
    PRIMARY KEY(id),
    FOREIGN KEY (log_detail_id) REFERENCES log_details(id)
);
