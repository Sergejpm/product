**Create tables:**

create table public.users
(
    id         integer generated always as identity
    constraint users_pk
    primary key,
    username   varchar(255)               not null,
    password   varchar(255)               not null,
    created_at timestamp(0) default now() not null
);

alter table public.users
    owner to postgres;

create table public.products
(
    id          integer generated always as identity
    constraint products_pk
    primary key,
    name        varchar(255)               not null
    constraint products_uk
    unique,
    description text                       not null,
    price       numeric(10, 2)             not null,
    created_at  timestamp(0) default now() not null
);

alter table public.products
    owner to postgres;

create table public.tokens
(
    id         integer generated always as identity
    constraint tokens_pk
    primary key,
    token      varchar(255)                                        not null,
    user_id    integer                                             not null
    constraint tokens_users_id_fk
    references public.users,
    created_at timestamp(0) default now()                          not null,
    expired_at timestamp(0) default (now() + '00:15:00'::interval) not null
);

alter table public.tokens
    owner to postgres;

create index tokens_user_id_index
    on public.tokens (user_id);


**Run service:**

docker-compose up --build -d

Postman collection: [product.postman_collection.json](product.postman_collection.json)