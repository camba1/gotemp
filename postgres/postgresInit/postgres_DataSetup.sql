-- Create the primary key generator

create sequence public.table_id_seq;

CREATE OR REPLACE FUNCTION public.next_id(OUT result bigint) AS $$
DECLARE
    our_epoch bigint := 1314220021721;
    seq_id bigint;
    now_millis bigint;
    shard_id int := 5;
BEGIN
    SELECT nextval('public.table_id_seq') % 1024 INTO seq_id;

    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
    result := (now_millis - our_epoch) << 23;
    result := result | (shard_id << 10);
    result := result | (seq_id);
END;
$$ LANGUAGE PLPGSQL;

-- Create table for promotions

CREATE TABLE "public"."promotion" (
    id  bigint not null default next_id(),
    name varchar(50) NOT NULL,
    description varchar(255),
    validfrom timestamp ,
    validthru timestamp ,
    active boolean DEFAULT false,
    customerid varchar(50)  NOT NULL,
    approvalstatus integer  DEFAULT 0,
    prevapprovalstatus integer DEFAULT 0,
    PRIMARY KEY (id)
);


-- Insert sample promotion data
insert into promotion (name, description, validfrom, validthru, customerid)
VALUES ('Promo1', 'Promo1 Desc', TIMESTAMP '2000-01-01 00:00:00', TIMESTAMP '2200-01-01 00:00:00','ducksrus'),
       ('Promo2', 'Promo2 Desc', TIMESTAMP '2000-01-01 00:00:00', TIMESTAMP '2200-01-01 00:00:00','patoloco'),
       ('Promo3', 'Promo3 Desc', TIMESTAMP '2000-01-01 00:00:00', TIMESTAMP '2200-01-01 00:00:00','ducksrus')