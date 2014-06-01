                                   Table "public.statistical_neighborhoods"
   Column   |          Type          |                                Modifiers                                
------------+------------------------+-------------------------------------------------------------------------
 gid        | integer                | not null default nextval('statistical_neighborhoods_gid_seq'::regclass)
 nbhd_id    | smallint               | 
 nbhd_name  | character varying(50)  | 
 typology   | character varying(33)  | 
 notes      | character varying(50)  | 
 shape_leng | numeric                | 
 shape_area | numeric                | 
 geom       | geometry(MultiPolygon) | 
Indexes:
    "statistical_neighborhoods_pkey" PRIMARY KEY, btree (gid)
