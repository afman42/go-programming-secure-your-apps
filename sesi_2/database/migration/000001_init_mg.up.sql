CREATE TYPE role_user AS ENUM ('admin', 'user');

CREATE TABLE users
(
    id serial primary key,
    email varchar(191) not null,
    password varchar(191) not null,
    role role_user not null,
    created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products
(
    id serial primary key,
    title varchar(191) not null,
    description varchar(191) not null,
    user_id serial not null,
    created_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

alter table products add foreign key (user_id) references users (id);

-- insert into users (email,password,role) values ("admin@example.com","$2a$08$P3VvD6ErYAoIpOfQ8D4IheoAZ0kGyzEcgz5G6OvmmFpkCJ47bhhhO","admin",now(),now());
-- insert into users (email,password,role) values ("user@example.com","$2a$08$P3VvD6ErYAoIpOfQ8D4IheoAZ0kGyzEcgz5G6OvmmFpkCJ47bhhhO","user",now(),now());