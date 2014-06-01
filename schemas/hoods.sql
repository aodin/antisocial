CREATE TABLE "hoods" (
  "id" INTEGER PRIMARY KEY,
  "name" VARCHAR,
  "population" INTEGER,
  "housing" INTEGER,
  "area" REAL,
  "crimes" INTEGER,
  "311" INTEGER,
  "foreclosures" INTEGER,
  "licenses" INTEGER,
  "score" REAL,
  "rank" INTEGER
);

COPY "hoods" FROM '/tmp/hoods.csv' WITH CSV HEADER;

ALTER TABLE "hoods" ADD COLUMN "geom" geometry('POLYGON', 4326);

Update hoods SET geom = statistical_neighborhoods.geo FROM statistical_neighborhoods WHERE statistical_neighborhoods.nbhd_id = id;

CREATE INDEX ON "hoods" USING GIST ("geom");

