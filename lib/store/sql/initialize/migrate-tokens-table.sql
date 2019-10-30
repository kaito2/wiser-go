CREATE TABLE tokens (
    id         INTEGER PRIMARY KEY,
    token      TEXT NOT NULL,
    docs_count INT NOT NULL,
    postings   BLOB NOT NULL
);

CREATE UNIQUE INDEX token_index ON tokens(token);