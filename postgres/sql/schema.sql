create table if not exists users
(
    id            uuid primary key,
    created_at    timestamptz not null,
    email         text unique not null,
    password_hash bytea       not null,
    activated     boolean     not null default false
);

create table if not exists tokens
(
    token_hash bytea primary key,
    user_id    uuid        not null references users (id) on delete cascade,
    scope      text        not null, -- TODO: enum
    expiry     timestamptz not null
);
