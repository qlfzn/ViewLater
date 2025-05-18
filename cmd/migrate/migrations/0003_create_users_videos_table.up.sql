CREATE TABLE user_videos (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    video_id INT REFERENCES videos(id) ON DELETE CASCADE,
    title TEXT,
    description TEXT,
    tags VARCHAR(100) [],
    created_at TIMESTAMPTZ DEFAULT now(),
    UNIQUE(user_id, video_id)
);
