create table users (
    id serial primary key,
    email varchar(30) unique not null,
    full_name varchar(30) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    deleted_at timestamp
);
create index idx__users__email on users(email);