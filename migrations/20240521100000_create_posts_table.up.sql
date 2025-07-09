CREATE TYPE post_visibility AS ENUM ('public', 'friends', 'private');

CREATE TABLE IF NOT EXISTS posts
(
    id             UUID PRIMARY KEY,
    user_id        UUID NOT NULL,
    video_url      TEXT NOT NULL,
    description    TEXT,
    duration       DOUBLE PRECISION NOT NULL,
    visibility     post_visibility NOT NULL DEFAULT 'public',
    hashtags       TEXT[],
    tags           TEXT[],
    likes_count    BIGINT NOT NULL DEFAULT 0,
    comments_count BIGINT NOT NULL DEFAULT 0,
    views_count    BIGINT NOT NULL DEFAULT 0,
    allow_comments BOOLEAN NOT NULL DEFAULT TRUE,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS posts_user_id_idx ON posts (user_id);

CREATE TRIGGER update_posts_updated_at
    BEFORE UPDATE
    ON
        posts
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column(); 