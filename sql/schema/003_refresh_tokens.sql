-- +goose Up
CREATE TABLE refresh_tokens (
    token TEXT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    user_id UUID REFERENCES users ON DELETE CASCADE,
    expires_at TIMESTAMP,
    revoked_at TIMESTAMP NULL
);

-- +goose Down
DROP TABLE refresh_tokens;