create table flats (
	flatId serial not null,
	houseId int not null,
	price int not null,
	rooms int not null,
	status text default 'created'::text not null
);


alter table flats
add constraint pk_flatId_houseId primary key(
	flatId,
	houseId
);


alter table flats
add constraint fk_houses_houseId foreign key(houseId)
references houses (houseId);


create index idx_houseId_searcher on flats (houseId);


comment on column flats.status is 'Enum: [created, approved, declined, on moderation]';
