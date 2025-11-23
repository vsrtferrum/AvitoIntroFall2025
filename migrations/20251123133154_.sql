-- +goose Up
-- +goose StatementBegin
CREATE TABLE pr (
    id VARCHAR(255) PRIMARY KEY, 
    name VARCHAR(255) not null,
    author INTEGER NOT NULL,
    opened BOOLEAN NOT NULL, 
    FOREIGN KEY (author) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pr IF EXISTS; 
-- +goose StatementEnd
