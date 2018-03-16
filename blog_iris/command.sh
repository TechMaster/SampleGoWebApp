create table posts (
    id serial primary key not null,
    title text,
    alias text,
    intro_text text,
    full_text text,
    image text,
    published text,
    published_at timestamptz,
    categories jsonb,
    type text null,
    created_at timestamptz,
    created_by text null,
    modified_at timestamptz,
    modified_by text null,
    author_visible text null
);

create table categories (
    id serial primary key not null,
    count int,
	name text,
	alias text,
	type text,
	description text
);
insert into categories (name) values ('linux'), ('react'), ('microservice'), ('html'), ('php'), ('java'), ('docker');


docker run --name db -e POSTGRES_PASSWORD=123 -d -p 5432:5432 postgres:latest

docker exec -it -u postgres db psql
