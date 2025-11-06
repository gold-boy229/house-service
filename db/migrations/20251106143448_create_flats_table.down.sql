drop index idx_houseId_searcher;

alter table flats drop constraint fk_houses_houseId;

alter table flats drop constraint pk_flatId_houseId;

drop table flats;
