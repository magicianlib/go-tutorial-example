create table user_info
(
    id      bigint auto_increment,
    name    varchar(32) null,
    deleted tinyint default 0 null,
    primary key (id)
);