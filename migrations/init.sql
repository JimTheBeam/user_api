CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null unique,
    created_at timestamp default current_timestamp,
    CONSTRAINT "pk_user_id" PRIMARY KEY (id)
);