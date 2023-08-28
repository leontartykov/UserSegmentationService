\c usersegmentappdb

alter table users
    add constraint c_users_pk primary key(id) not null;

alter table segments
    add constraint c_segments_pk primary key(id) not null;

alter table usersSegments
    add constraint c_users_fk foreign key (user) references users(nickname)
    add constraint c_segments_fk foreign key (segment) references segments(name);
