CREATE TABLE domain
(
    id serial not null  unique,
    name varchar(255) not null unique,
    status smallint not null
);

CREATE TABLE labels
(
    id serial not null  unique,
    domain_id int not null,
    metric_name varchar(255) not null unique,
    metric_id int not null
);

CREATE TABLE counters
(
    id serial not null  unique,
    metric_name varchar(255) not null unique,
    metric_id int not null unique,
    label_id int not null,
    created_at date default current_date
);