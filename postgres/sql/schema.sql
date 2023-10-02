create table users
(
    id            uuid primary key,
    email         text unique not null,
    password_hash bytea       not null
);

create table sessions
(
    id         uuid primary key,
    user_id    text         not null references users (id) on delete cascade,
    token_hash bytea unique not null
);
