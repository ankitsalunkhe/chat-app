ALTER TABLE chat_app.credentials
    ADD COLUMN is_google_signup BOOLEAN DEFAULT FALSE;
ALTER TABLE chat_app.credentials
    ALTER COLUMN password DROP NOT NULL;