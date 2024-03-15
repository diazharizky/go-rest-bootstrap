CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(30) UNIQUE NOT NULL,
    full_name VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);
