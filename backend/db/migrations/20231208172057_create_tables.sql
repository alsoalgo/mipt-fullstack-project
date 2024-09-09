-- +goose Up
-- +goose StatementBegin
create table users (
    id              bigserial               primary key,
    email           text        not null,
    password_hash   text        not null,
    user_role       text        not null    default 'default',
    first_name      text,
    last_name       text,
    sur_name        text,
    created_at      timestamptz not null    default now(),
    updated_at      timestamptz not null    default now()
);

create table hotel_order (
    id          bigserial               primary key,
    user_id     bigint      not null,
    hotel_id    bigint      not null,
    date_from   timestamptz not null,
    date_to     timestamptz not null,
    first_name  text        not null,
    last_name   text        not null,
    sur_name    text        not null,
    created_at  timestamptz not null    default now(),
    updated_at  timestamptz not null    default now()
);

create table hotel (
    id          bigserial               primary key,
    city        text        not null,
    title       text        not null,
    description text        not null,
    image_url   text        not null,
    created_at  timestamptz not null    default now(),
    updated_at  timestamptz not null    default now()
);

create table user_question (
    id          bigserial               primary key,
    user_id     bigint      not null,
    title       text        not null,
    question    text        not null,
    status      text        not null    default 'unresolved',
    created_at  timestamptz not null    default now(),
    updated_at  timestamptz not null    default now()
);

create table tokens (
    id          bigserial                   primary key,
    user_id     bigint          not null,
    token       text            not null,
    created_at  timestamptz     not null    default now(),
    updated_at  timestamptz     not null    default now(),
    expires_at  timestamptz     not null    default now() + '1 hour'::interval
);

create table available_room (
    id              bigserial               primary key,
    hotel_id        bigint      not null,
    room_count      integer     not null    default 0,
    available_date  date        not null,
    created_at      timestamptz not null    default now(),
    updated_at      timestamptz not null    default now()
);

create table popular_destination (
    id          bigserial               primary key,
    city        text        not null,
    cost        integer     not null,
    image_url   text        not null,
    created_at  timestamptz not null    default now(),
    updated_at  timestamptz not null    default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
drop table if exists hotel_order;
drop table if exists hotel;
drop table if exists user_question;
drop table if exists tokens;
drop table if exists available_room;
drop table if exists popular_destination;
-- +goose StatementEnd
