CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    author_id INT,
    title VARCHAR(50) NOT NULL,
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY(author_id) REFERENCES users(id)
);
