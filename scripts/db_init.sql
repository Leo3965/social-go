INSERT INTO users (email, username, password)
VALUES ('leo@example.com',
        'leo',
        convert_to('password123', 'UTF8'));
