

create table if not exists todos (
    todo_id varchar(255) primary key,
    title varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    deleted_at datetime
);
