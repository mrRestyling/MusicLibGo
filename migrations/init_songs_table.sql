CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    group_id INTEGER NOT NULL,
    release_date TEXT NOT NULL,
    text TEXT NOT NULL,
    link TEXT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE INDEX idx_songs_group_id ON songs (group_id);
CREATE INDEX idx_groups_name ON groups (name);