package model

import (
	"database/sql"
	"time"
)

type House struct {
	Id        int            // not NULL
	Address   string         // not NULL
	Year      int            // not NULL
	Developer sql.NullString // can be NULL
	CreatedAt time.Time      // has default value
	UpdatedAt time.Time      // has default value
}

// table houses (
// 	  houseId serial primary key,
// 	  address text not null,
// 	  year int not null,
// 	  developer text null,
// 	  created_at TIMESTAMPTZ default now(),
// 	  updated_at TIMESTAMPTZ default now()
// )
