create table users
(
    id           bigserial primary key,
    full_name    varchar(255),
    email        varchar(255) not null unique,
    password     varchar(255),
    number_phone varchar(100),
    active_phone boolean                  default false,
    active_email boolean                  default false,
    role         bigint                   default 2,
    status_id    bigint,
    created_at   timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at   timestamp with time zone default CURRENT_TIMESTAMP
);


create table categories
(
    id         bigserial primary key,
    name       varchar(255),
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

create table subcategories
(
    id            bigserial primary key,
    name          varchar(255),
    categories_id bigint,
    created_at    timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at    timestamp with time zone default CURRENT_TIMESTAMP
);

create table gender
(
    id           bigserial primary key,
    name         varchar(255),
    categoriesid bigint,
    created_at   timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at   timestamp with time zone default CURRENT_TIMESTAMP
);

create table brands
(
    id            bigserial primary key,
    name          varchar(255),
    categories_id bigint,
    created_at    timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at    timestamp with time zone default CURRENT_TIMESTAMP
);

create table species
(
    id               bigserial primary key,
    name             varchar(255),
    subcategories_id bigint,
    brands_id        bigint,
    created_at       timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at       timestamp with time zone default CURRENT_TIMESTAMP
);

create table products
(
    id               bigserial primary key,
    name             varchar(255),
    price            float8,
    old_price        float8,
    color_id         bigint,
    size_id          bigint,
    material_id      bigint,
    season           varchar(20),
    species_id       bigint,
    inventory_number bigint,
    brands_id        bigint,
    gender_id        bigint,
    status_id        bigint,
    created_at       timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at       timestamp with time zone default CURRENT_TIMESTAMP
);

create table colors
(
    id         bigserial primary key,
    name       varchar(255),
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

create table roles
(
    id         bigserial primary key,
    name       varchar(255),
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

create table size
(
    id         bigserial primary key,
    name       varchar(255),
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

create table materials
(
    id         bigserial primary key,
    name       varchar(255),
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP
);

create table status
(
    id         bigserial primary key,
    name       varchar(255),
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP
);
