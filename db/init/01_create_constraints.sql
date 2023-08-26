\c usersegmentappdb

alter table users
    add constraint c_users_pk primary key(id) not null;

alter table segments
    add constraint c_segments_pk primary key(id) not null;

alter table usersSegments
    add constraint c_users_fk foreign key (user_id) references users(id)
    add constraint c_segments_fk foreign key (segment_id) references segments(id);
