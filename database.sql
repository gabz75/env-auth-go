CREATE EXTENSION citext;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id serial primary key,
  email citext UNIQUE,
  password varchar(60)
);

DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions (
  user_id serial references users(id),
  token varchar(60)
);
