CREATE SCHEMA IF NOT EXISTS chat_app;
CREATE TABLE chat_app.credentials (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);