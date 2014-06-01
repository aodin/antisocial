1. Create schemas

    * Crimes
    * 311

2. Clean Data

    "Wood burning on "Red Days"" > "Wood burning on 'Red Days'"
    "Request for Proclamation, "Paralegal Day"" > "Request for Proclamation, 'Paralegal Day'"
    "How did "911"" > "How did '911'"


3. Copy in event data

    # TODO remove header from second file
    (cat 311_service_data_2013.csv && cat 311_service_requests_2014.csv) > 311.csv

    COPY "crimes" FROM '/tmp/crime.csv' WITH CSV HEADER;
    COPY "311" FROM '/tmp/311.csv' WITH CSV HEADER;
    COPY "foreclosures" FROM '/tmp/forclosures.csv' WITH CSV HEADER;
    COPY "licenses" FROM '/tmp/licenses.csv' WITH CSV HEADER;


4. Load in neighborhood data

    CREATE EXTENSION postgis;
    CREATE EXTENSION postgis_topology;

    shp2pgsql /tmp/statistical_neighborhoods.shp > neighborhoods.sql

    sudo -u postgres psql -d antisocial < /tmp/neighborhoods.sql

    Convert the multi-polygon to a polygon

    ALTER TABLE "statistical_neighborhoods" ADD COLUMN "geo" geometry('POLYGON', 4326);

    UPDATE "statistical_neighborhoods" SET "geo" = st_geometryn(geom, 1);

    CREATE INDEX ON "statistical_neighborhoods" USING GIST ("geo");
    
    Output GeoJSON for a neighborhood
    
     SELECT ST_AsGeoJSON("geo") FROM "statistical_neighborhoods";


5. Alter the schema for PostGIS

    ALTER TABLE "crimes" ADD COLUMN "loc" geometry('POINT', 4326);
    UPDATE "crimes" SET "loc" = ST_GeometryFromText('POINT(' || "longitude" || ' ' || "latitude" || ')', 4326);

    ALTER TABLE "foreclosures" ADD COLUMN "loc" geometry('POINT', 4326);
    UPDATE "foreclosures" SET "loc" = ST_GeometryFromText('POINT(' || "longitude" || ' ' || "latitude" || ')', 4326);

    ALTER TABLE "licenses" ADD COLUMN "loc" geometry('POINT', 4326);
    UPDATE "licenses" SET "loc" = ST_GeometryFromText('POINT(' || "longitude" || ' ' || "latitude" || ')', 4326);

Test

    VACUUM ANALYZE SELECT "nbhd_name", ST_Area(ST_Transform(geom,4326)::geography) FROM "statistical_neighborhoods";


311 events

    ALTER TABLE "simple311" ADD COLUMN "loc" geometry('POINT', 4326);

    UPDATE "simple311" SET "loc" = ST_GeometryFromText('POINT(' || "longitude" || ' ' || "latitude" || ')', 4326);



6. Queries

    Don't forget to average by area! Area of geometry:

```sql

SELECT UpdateGeometrySRID('statistical_neighborhoods', 'geom', 4326);

SELECT "nbhd_name", ST_Area("geom") FROM "statistical_neighborhoods";

SELECT "nbhd_name", ST_Area(ST_Transform(geom,4326)::geography) FROM "statistical_neighborhoods";
```

    Crimes by neighborhoods
```sql

SELECT COUNT("crimes"."id"), "statistical_neighborhoods"."nbhd_name" FROM "crimes", "statistical_neighborhoods" where ST_Intersects(ST_GeometryFromText('POINT(' || "longitude" || ' ' || "latitude" || ')', 4326), ST_SetSRID("geom", 4326)) GROUP BY "statistical_neighborhoods"."nbhd_name";

-- After indexing

SELECT COUNT("crimes"."id"), "statistical_neighborhoods"."nbhd_name" FROM "crimes", "statistical_neighborhoods" where ST_Intersects("loc", "geom") GROUP BY "statistical_neighborhoods"."nbhd_name";


```

    311 events by neighborhood
```sql

SELECT COUNT("simple311"."summary"), "statistical_neighborhoods"."nbhd_name" FROM "simple311", "statistical_neighborhoods" where ST_Intersects("loc", "geom") GROUP BY "statistical_neighborhoods"."nbhd_name" ORDER BY "statistical_neighborhoods"."nbhd_name" ASC;

```

