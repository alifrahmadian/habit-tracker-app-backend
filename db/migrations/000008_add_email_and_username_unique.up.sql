ALTER TABLE users
ADD COLUMN email VARCHAR(255),
ADD CONSTRAINT users_email_unique UNIQUE (email);

ALTER TABLE users
ADD CONSTRAINT users_username_unique UNIQUE (username);