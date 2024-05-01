create table articles (
    id serial primary key,
    author_id bigint not null,
    title varchar(100) not null,
    content text,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    deleted_at timestamp,
    constraint fk__user foreign key (author_id) references users(id)
);
create index idx__articles__title on articles(title);