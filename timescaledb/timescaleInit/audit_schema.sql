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

CREATE TABLE "public"."audit" (
                                  actiontime timestamp NOT NULL,     --Time the action took place in the original service
--                                id  bigint not null default next_id(),
                                  topic varchar(50) NOT NULL,        --Topic to which the change was published
                                  service varchar(50) NOT NULL,      -- Name of service that caused the modification
                                  actionFunc varchar(255) NOT NULL,  --Function in the service that caused the modification
                                  actionType varchar(50) NOT NULL,   --Indicates what action was performed (Insert, delete, etc)
                                  objectName varchar(250) NOT NULL,  --Name of object (table) to which the modified record belongs
                                  objectId varchar(50)  NOT NULL,         --Id of the record modified
                                  performedBy bigint NOT NULL ,      --Id of user that did the change
                                  recordedtime timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,   --Time we recorded action in audit table
                                  objectDetail bytea NOT NULL,        -- Information of the record modified
                                  id  bigint not null default next_id()
);

CREATE INDEX ON "public"."audit" (objectName, objectId, actiontime DESC);
CREATE INDEX ON "public"."audit" (id, actiontime DESC);

SELECT create_hypertable('audit', 'actiontime');

-- Insert sample  data
insert into audit (actiontime, topic, service, actionFunc, actionType, objectName, objectId, performedBy, objectDetail )
VALUES (TIMESTAMP '2000-01-01 00:00:00', 'audit', 'user', 'afterCreateUser','insert', 'user', '123456789', '2345678', '\\\\'::bytea),
       (TIMESTAMP '2000-02-01 00:00:00', 'audit', 'user', 'afterDeleteUser','delete', 'user', '12345678910', '2345678', '\\\\'::bytea),
       (TIMESTAMP '2000-03-01 00:00:00', 'audit', 'user', 'afterUpdateUser','update', 'user', '12345678911', '2345678', '\\\\'::bytea);
