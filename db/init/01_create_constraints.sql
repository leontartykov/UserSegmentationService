\c usersegmentappdb

alter table users
    add constraint c_users_pk primary key(id),
    add constraint c_users_nickname_unique UNIQUE(nickname);

alter table segments
    add constraint c_segments_pk primary key(id),
    add constraint c_segments_name_unique UNIQUE(name);

alter table users_segments
    add constraint c_segments_fk foreign key (segment_name) references segments(name);
