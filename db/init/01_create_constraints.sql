\c usersegmentappdb

alter table users
    add constraint c_users_pk primary key(id),
    add constraint c_users_nickname_unique UNIQUE(nickname);

alter table segments
    add constraint c_segments_pk primary key(id),
    add constraint c_segments_name_unique UNIQUE(name);

alter table usersSegments
    add constraint c_users_fk foreign key (userName) references users(nickname),
    add constraint c_segments_fk foreign key (segmentName) references segments(name);
