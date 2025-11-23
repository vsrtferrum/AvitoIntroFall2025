-- +goose Up
-- +goose StatementBegin

CREATE INDEX IF NOT EXISTS idx_users_is_active ON users(is_active);

CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

CREATE INDEX IF NOT EXISTS idx_teams_user_id ON teams(user_id);

CREATE INDEX IF NOT EXISTS idx_pr_author ON pr(author);

CREATE INDEX IF NOT EXISTS idx_pr_opened ON pr(opened);

CREATE INDEX IF NOT EXISTS idx_review_reviewer ON review(reviewer);


CREATE INDEX IF NOT EXISTS idx_review_id ON review(id);

CREATE INDEX IF NOT EXISTS idx_teams_name_user_id ON teams(name, user_id);

CREATE INDEX IF NOT EXISTS idx_review_stats ON review(reviewer, id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_is_active;
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_teams_user_id;
DROP INDEX IF EXISTS idx_pr_author;
DROP INDEX IF EXISTS idx_pr_opened;
DROP INDEX IF EXISTS idx_review_reviewer;
DROP INDEX IF EXISTS idx_review_id;
DROP INDEX IF EXISTS idx_teams_name_user_id;
DROP INDEX IF EXISTS idx_review_stats;
-- +goose StatementEnd