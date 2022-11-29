CREATE TABLE domain
(
    id serial not null  unique,
    name varchar(255) not null unique,
    status smallint not null
);

CREATE TABLE labels
(
    id serial not null  unique,
    domain_id varchar(255) not null,
    metric_name varchar(255) not null unique,
    metric_id int not null
);

CREATE TABLE counters
(
    id serial not null  unique,
    metric_name varchar(255) not null unique,
    metric_id int not null,
    label_id int
);