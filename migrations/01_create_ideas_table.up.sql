create table ideas (
    id serial,
    user_id int,
    title text,
    description text,
    created_at timestamp with time zone default now(),
    PRIMARY KEY (title, user_id)
);
