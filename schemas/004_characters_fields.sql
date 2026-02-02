-- +goose Up
ALTER TABLE characters
    ADD COLUMN archetype TEXT,
    ADD COLUMN campaign TEXT,
    ADD COLUMN game TEXT;

-- +goose Down
ALTER TABLE characters
    DROP COLUMN game,
    DROP COLUMN campaign,
    DROP COLUMN archetype;
