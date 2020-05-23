-- Create 'users' table
create table if not exists public.players (
    id serial primary key,
    name varchar(255) not null,
    score integer not null
);