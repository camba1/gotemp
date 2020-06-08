-- topic:        topic,
-- 		objectToSend: objectToSend,
-- 		header: auditMsgHeader{
-- 			"service":		serviceName,
-- 			"actionFunc":  actionFunc,
-- 			"actionType":  actionType,
-- 			"objectId":    string(objectId),
-- 			"performedBy": string(performedBy),
-- 		},


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
                                  time timestamp NOT NULL,
--                                       id  bigint not null default next_id(),
                                  topic varchar(50) NOT NULL,
                                  service varchar(50) NOT NULL,
                                  actionFunc varchar(255) NOT NULL,
                                  actionType varchar(50) NOT NULL,
                                  objectName varchar(250) NOT NULL,
                                  objectId bigint  NOT NULL,
                                  performedBy bigint NOT NULL ,
                                  objectDetail bytea NOT NULL
);

CREATE INDEX ON "public"."audit" (objectName, objectId, time DESC);

SELECT create_hypertable('audit', 'time');

-- Insert sample  data
insert into audit (time, topic, service, actionFunc, actionType, objectName, objectId, performedBy, objectDetail )
VALUES (TIMESTAMP '2000-01-01 00:00:00', 'audit', 'user', 'afterCreateUser','insert', 'user', '123456789', '2345678', '\\\\'::bytea),
       (TIMESTAMP '2000-02-01 00:00:00', 'audit', 'user', 'afterDeleteUser','delete', 'user', '12345678910', '2345678', '\\\\'::bytea),
       (TIMESTAMP '2000-03-01 00:00:00', 'audit', 'user', 'afterUpdateUser','update', 'user', '12345678911', '2345678', '\\\\'::bytea);
