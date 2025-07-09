DROP TRIGGER IF EXISTS update_posts_updated_at ON posts;

DROP INDEX IF EXISTS posts_user_id_idx;

DROP TABLE IF EXISTS posts;
 
DROP TYPE IF EXISTS post_visibility; 