alter table houses 
add constraint unique_address_year_developer_idx unique(address, year, developer);