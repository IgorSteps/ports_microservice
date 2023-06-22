package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PortSchema struct {
	Id          uuid.UUID       `gorm:"primaryKey; type:uuid; column=id;"`
	Name        string          `gorm:"type:string; column=name;"`
	City        string          `gorm:"type:string; column=city;"`
	Country     string          `gorm:"type:string; column=country;"`
	Alias       pq.StringArray  `gorm:"type:text[]; column=alias;"`
	Regions     pq.StringArray  `gorm:"type:text[]; column=regions;"`
	Coordinates pq.Float32Array `gorm:"type:float[]; column=coordinates;"`
	Province    string          `gorm:"type:string; column=province;"`
	Timezone    string          `gorm:"type:string; column=timezone;"`
	Unlocs      pq.StringArray  `gorm:"type:text[]; column=unlocs;"`
	Code        string          `gorm:"type:string; column=code;"`
}
