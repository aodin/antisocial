 HashAggregate  (cost=4705834.57..4705835.35 rows=78 width=24)
   ->  Nested Loop  (cost=0.00..4705805.06 rows=5902 width=24)
         Join Filter: ((crimes.loc && statistical_neighborhoods.geom) AND _st_intersects(crimes.loc, statistical_neighborhoods.geom))
         ->  Seq Scan on crimes  (cost=0.00..13568.06 rows=227006 width=45)
         ->  Materialize  (cost=0.00..23.17 rows=78 width=43)
               ->  Seq Scan on statistical_neighborhoods  (cost=0.00..22.78 rows=78 width=43)
(6 rows)


CREATE INDEX ON "statistical_neighborhoods" USING GIST ("geom");
CREATE INDEX ON "statistical_neighborhoods" USING GIST ("geo");

CREATE INDEX ON "crimes" USING GIST ("loc");

CREATE INDEX ON "simple311" USING GIST ("loc");

CREATE INDEX ON "licenses" USING GIST ("loc");


 HashAggregate  (cost=7518.47..7519.25 rows=78 width=24)
   ->  Nested Loop  (cost=0.00..7488.96 rows=5902 width=24)
         ->  Seq Scan on statistical_neighborhoods  (cost=0.00..22.78 rows=78 width=43)
         ->  Index Scan using crimes_loc_idx on crimes  (cost=0.00..95.64 rows=8 width=45)
               Index Cond: (loc && statistical_neighborhoods.geom)
               Filter: _st_intersects(loc, statistical_neighborhoods.geom)
(6 rows)