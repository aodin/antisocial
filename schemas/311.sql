CREATE TABLE "311" (
    "summary" varchar,
    "status" varchar,
    "source" varchar,
    "created_date" date,
    "created" timestamp,
    "closed_date" date,
    "closed" timestamp,
    "resolution" varchar,
    "customer_zip" varchar,
    "address" varchar,
    "address_cont" varchar,
    "intersection" varchar,
    "intersection_cont" varchar,
    "zip" varchar,
    "longitude" real,
    "latitude" real,
    "agency" varchar,
    "division" varchar,
    "area" varchar,
    "type" varchar,
    "topic" varchar,
    "council_district" varchar,
    "police_district" varchar,
    "neighborhood" varchar
);

CREATE TABLE "simple311" (
    "summary" varchar,
    "created" timestamp,
    "address" varchar,
    "zip" varchar,
    "longitude" real,
    "latitude" real,
    "type" varchar,
    "topic" varchar,
    "neighborhood" varchar
);