-- +goose Up
-- +goose StatementBegin
CREATE TABLE review(
    id VARCHAR(255) NOT NULL, 
    reviewer VARCHAR(255) NOT NULL, 
    FOREIGN KEY (id) REFERENCES pr(id), 
    FOREIGN KEY (reviewer) REFERENCES users(id)
    CONSTRAINT one_user_to_one_pr UNIQUE (id, reviewer)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS review;
-- +goose StatementEnd
