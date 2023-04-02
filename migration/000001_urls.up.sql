CREATE TABLE IF NOT EXISTS urls (
    user_id varchar(255) not null,
    short_url varchar(255) not null,
    original_url varchar(255) not null unique,
    uses integer not null
);