ALTER TABLE chat_app.credentials
    DROP COLUMN is_google_signup;
ALTER TABLE chat_app.credentials
    ALTER COLUMN password ADD NOT NULL;