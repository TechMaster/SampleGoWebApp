# A Blog with Iris

### Preview Published Posts

![](https://raw.githubusercontent.com/TechMaster/SampleGoWebApp/master/blog_iris/images/preview.jpg)

### All post

![](https://raw.githubusercontent.com/TechMaster/SampleGoWebApp/master/blog_iris/images/allposts.jpg)

### Details post

![](https://raw.githubusercontent.com/TechMaster/SampleGoWebApp/master/blog_iris/images/details.jpg)

## Install


# create table posts with this form
    
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

### then
    drop table categories;
    create table categories (
    id serial primary key not null,
    count int,
	name text,
	alias text,
	type text,
	description text
    );
    insert into categories (name) values ('linux'), ('react'), ('microservice'), ('html'), ('php'), ('java'), ('docker');

### and run command

    go run main.go