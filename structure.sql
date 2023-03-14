create table users (
  id uuid default uuid_generate_v4 (),
  name varchar not null,
  email varchar not null,
  passhash varchar not null,
  balance numeric (10, 2)
);