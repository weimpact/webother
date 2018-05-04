create table files (
    id serial,
    user_id int,
    location text,
    created_at timestamp with time zone,
    PRIMARY KEY (user_id, location)
)
