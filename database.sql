DROP DATABASE IF EXISTS go_auth;
CREATE DATABASE go_auth;

\c go_auth;

CREATE EXTENSION citext;

CREATE TABLE users (
  id serial primary key,
  email citext UNIQUE,
  password varchar(60)
);

CREATE TABLE sessions (
  id serial primary key,
  user_id serial references users(id),
  token varchar(233)
);
