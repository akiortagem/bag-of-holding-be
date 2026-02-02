-- +goose Up
CREATE TABLE characters (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS characters_user_id_idx ON characters (user_id);

-- +goose Down
DROP INDEX IF EXISTS characters_user_id_idx;
DROP TABLE characters;
