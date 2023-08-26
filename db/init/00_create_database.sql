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

create table usersSegments (
    user_id int,
    segment_id int,
    added_at date,
    deleted_at date,
    expired_at date
);
