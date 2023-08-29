create database usersegmentappdb;

\c usersegmentappdb

create table users (
    id serial,
    nickname text
);

create table segments (
    id serial,
    name text
);

create table users_segments (
    user_id integer,
    segment_name text,
    added_at date,
    deleted_at date,
    expired_at date
);
