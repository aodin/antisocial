Bore Score
==========


Neighborhood Data
-----------------

Size
Population
Crimes / 1000 people (per year?)
311 events / 1000 people (per year?)
Bars / km2
Foreclosures / dwelling



### Queries

Building the neighborhood area

```sql
SELECT "nbhd_id", ST_Area("geom") FROM "statistical_neighborhoods" ORDER BY "nbhd_id" ASC;

SELECT "nbhd_id", ST_Area(ST_Transform(geom,4326)::geography) FROM "statistical_neighborhoods" ORDER BY "nbhd_id" ASC;
```

Crimes in 2013 by neighborhood

```sql
SELECT COUNT("crimes"."id"), "statistical_neighborhoods"."nbhd_id" FROM "crimes", "statistical_neighborhoods" WHERE ST_Intersects("loc", "geom") AND date_part('year', "crimes"."reported") = 2013 GROUP BY "statistical_neighborhoods"."nbhd_id" ORDER BY "statistical_neighborhoods"."nbhd_id" ASC;
```

311 Calls in 2013 by neighborhood

```sql
SELECT COUNT("simple311"."summary"), "statistical_neighborhoods"."nbhd_id" FROM "simple311", "statistical_neighborhoods" WHERE ST_Intersects("loc", "geom") AND date_part('year', "simple311"."created") = 2013 GROUP BY "statistical_neighborhoods"."nbhd_id" ORDER BY "statistical_neighborhoods"."nbhd_id" ASC;
```

Foreclosures in 2012 by neighborhood

```sql
SELECT COUNT("foreclosures"."filenumber"), "statistical_neighborhoods"."nbhd_id" FROM "foreclosures", "statistical_neighborhoods" WHERE ST_Intersects("loc", "geom") AND "foreclosures"."year" = 2012 GROUP BY "statistical_neighborhoods"."nbhd_id" ORDER BY "statistical_neighborhoods"."nbhd_id" ASC;
```

Active liquor licenses by neighborhood

```sql
SELECT COUNT("licenses"."id"), "statistical_neighborhoods"."nbhd_id" FROM "licenses", "statistical_neighborhoods" WHERE ST_Intersects("loc", "geom") AND "licenses"."expires" > now() GROUP BY "statistical_neighborhoods"."nbhd_id" ORDER BY "statistical_neighborhoods"."nbhd_id" ASC;
```
