CREATE TABLE users
(
    id         bigserial,
    username   varchar NOT NULL,
    full_name  varchar NOT NULL,
    email      varchar NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    CONSTRAINT users_pk PRIMARY KEY (id),
    CONSTRAINT users_unq_username UNIQUE (username),
    CONSTRAINT users_unq_email UNIQUE (email)
);

CREATE INDEX users_idx_full_name ON users USING btree (full_name);
