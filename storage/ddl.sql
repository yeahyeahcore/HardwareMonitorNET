CREATE OR REPLACE FUNCTION insert_time ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.created_at = now();
    RETURN NEW;
END;
$$
LANGUAGE 'plpgsql';

DROP TABLE IF EXISTS devices;
CREATE TABLE devices (
    id varchar(20) NOT NULL PRIMARY KEY,
    pc_name varchar(128) NOT NULL,
    mac_address varchar(17) NOT NULL,
    cpu_name varchar(128),
    hdd_name varchar(128),
    gpu_name varchar(128),
    gpu_memory real
);

DROP IF EXISTS TABLE parameters;
CREATE TABLE parameters (
    id bigserial NOT NULL PRIMARY KEY,
    device_id varchar(20) NOT NULL,
    created_at timestamp without time zone,
    cpu_temp smallint[],
    cpu_clock real[],
    cpu_load smallint[],
    power_cpu_package real,
    power_cpu_cores real,
    power_cpu_graphics real,
    power_cpu_dram real,
    memory_load real,
    memory_used real,
    memory_available real,
    hdd_temp smallint,
    gpu_clocks real,
    gpu_load real,
    gpu_memory_used real,
    gpu_memory_free real
);

CREATE INDEX ON parameters (device_id);
CREATE INDEX ON parameters (created_at);

CREATE TRIGGER on_insert
    BEFORE INSERT ON parameters
    FOR EACH ROW
    EXECUTE PROCEDURE insert_time ();