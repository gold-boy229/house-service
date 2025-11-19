create EXTENSION if not exists "uuid-ossp";

create table app_users (
	userId uuid default gen_random_uuid() primary key,
	role text default 'client'::text,
	email text not null,
	password_hash text not null
);


alter table app_users
add constraint unique_role_email unique (role, email);

comment on column app_users.role is 'Enum: [client, moderator]';
