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

-- CREATE TABLE songs (
--     id SERIAL PRIMARY KEY,
--     title VARCHAR(255) NOT NULL,
--     group_id INTEGER NOT NULL,
--     release_date TEXT NOT NULL,
--     text TEXT NOT NULL,
--     link TEXT NOT NULL,
--     FOREIGN KEY (group_id) REFERENCES groups(id),
--     UNIQUE (title, group_id)
-- );