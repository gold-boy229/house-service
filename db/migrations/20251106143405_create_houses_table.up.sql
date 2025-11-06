create table houses (
    houseId serial primary key,
    address text not null,
    year int not null,
    developer text null,
    created_at TIMESTAMPTZ default now(),
    updated_at TIMESTAMPTZ default now()
);


comment on column houses.year is 'Year of construction';
comment on column houses.updated_at is 'Datetime of last flat creation in this house';
