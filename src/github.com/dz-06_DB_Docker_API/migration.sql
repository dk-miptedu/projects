--drop table if exists conn_log;
create table if not exists conn_log (user_id bigint, ip_addr varchar(15), ts timestamp);
create index idx_user_id on conn_log (user_id);
insert into conn_log (user_id, ip_addr, ts)
select (random()*6 +1)::int as user_id,
       ('{"127.0.0.1","127.0.0.2","127.0.0.3","127.0.0.4"}'::text[])[floor(random()*4)+1] as ip_addr,
       timestamp '2024-01-01 00:00:00' + (random()*(365*24*60*60))* interval '1 second'
from generate_series (1, 1000) s(i);