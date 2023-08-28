create database usersegmentappdb;

\c usersegmentappdb

create table users (
    id serial,
    nickname text UNIQUE
);

create table segments (
    id serial,
    name text UNIQUE
);

create table usersSegments (
    userName text,
    segmentName text,
    added_at date,
    deleted_at date,
    expired_at date
);
