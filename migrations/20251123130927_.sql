-- +goose Up
-- +goose StatementBegin
CREATE TABLE teams(
    name VARCHAR(255) UNIQUE,
    user_id VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
    CONSTRAINT one_user_to_one_team UNIQUE (name, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE teams if EXISTS;
-- +goose StatementEnd
