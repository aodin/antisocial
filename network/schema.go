package network

import (
	. "github.com/aodin/aspect"
	"github.com/aodin/aspect/postgis"
)

var Hoods = Table("hoods",
	Column("id", Integer{PrimaryKey: true}),
	Column("name", String{}),
	Column("population", Integer{}),
	Column("housing", Integer{}),
	Column("area", Real{}),
	Column("crimes", Integer{}),
	Column("311", Integer{}),
	Column("foreclosures", Integer{}),
	Column("licenses", Integer{}),
	Column("score", Real{}),
	Column("rank", Integer{}),
	Column("geom", postgis.GeometryPolygon{4326}),
)

type Hood struct {
	Id           int64   `json:"id" db:"id"`
	Name         string  `json:"name" db:"name"`
	Population   int64   `json:"population" db:"population"`
	Housing      int64   `json:"housing" db:"housing"`
	Area         float64 `json:"area" db:"area"`
	Crimes       int64   `json:"crimes" db:"crimes"`
	Calls311     int64   `json:"calls" db:"311"`
	Foreclosures int64   `json:"foreclosures" db:"foreclosures"`
	Licenses     int64   `json:"licenses" db:"licenses"`
	Score        float64 `json:"score" db:"score"`
	Rank         int64   `json:"rank" db:"rank"`
	Geom         string  `json:"geom" db:"geom"`
}

func (h Hood) String() string {
	return h.Name
}
