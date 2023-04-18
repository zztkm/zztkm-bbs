DROP TABLE IF EXISTS posts;
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at INTEGER NOT NULL
);
CREATE INDEX idx_posts_on_created_at ON posts (posts DESC);
INSERT INTO posts (name, body, created_at) VALUES (
    'zztkm',
	'初投稿！',
    unixepoch()
);
