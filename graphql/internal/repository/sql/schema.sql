CREATE TABLE users (
    id text PRIMARY KEY NOT NULL,
    name text NOT NULL
);

CREATE TABLE todos (
    id text PRIMARY KEY NOT NULL,
    user_id text NOT NULL,
    text text NOT NULL,
    done boolean NOT NULL DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

