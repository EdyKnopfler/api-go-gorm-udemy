CREATE TABLE IF NOT EXISTS candidates (
    id BIGSERIAL primary key,
    name TEXT not null,
    created_at TIMESTAMP default now()
);
