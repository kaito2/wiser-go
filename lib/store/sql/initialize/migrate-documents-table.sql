CREATE TABLE documents (
    id      INTEGER PRIMARY KEY,
    title   TEXT NOT NULL,
    body    TEXT NOT NULL
);

CREATE UNIQUE INDEX title_index ON documents(title);