create table `user`
(
    `id`          bigint(20)                             not null auto_increment,
    `user_id`     bigint(20)                             not null,
    -- utf8mb4_general_ci 这个就是排序的时候有点速度上的优势 仅此而已
    `username`    varchar(64) collate utf8mb4_general_ci not null,
    `password`    varchar(64) collate utf8mb4_general_ci not null,
    `email`       varchar(64) collate utf8mb4_general_ci,
    `gender`      tinyint(4)                             not null default '0',
    -- 这个就是插入数据的时候 可以自动设置时间
    `create_time` timestamp                              not null default current_timestamp,
    -- 更新数据饿的时候 时间戳也可以自动更新
    `update_time` timestamp                              not null default current_timestamp on update current_timestamp,
    primary key (`id`),
    -- 建立索引 加快存储速度 同时unique key 代表是唯一的 不能重复
    unique key `idx_username` (`username`) using btree,
    unique key `idx_user_id` (`user_id`) using btree
) default charset = utf8mb4
  collate = utf8mb4_general_ci;


Drop table if exists `community`;
create table `community`
(
    `id`             int(11)                                 not null auto_increment,
    -- 板块id
    `community_id`   int(10) unsigned                        not null,
    -- 板块名称
    `community_name` varchar(128) collate utf8mb4_general_ci not null,
    -- 板块介绍
    `introduction`   varchar(256) collate utf8mb4_general_ci not null,
    `create_time`    timestamp                               not null default current_timestamp,
    `update_time`    timestamp                               not null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci;

insert into community (community_id, community_name, introduction) values ('1', 'Go', 'Go语言');
insert into community (community_id, community_name, introduction) values ('2', '狮子头', '扬州菜');
insert into community (community_id, community_name, introduction) values ('3', '铁拳', '经典游戏');
insert into community (community_id, community_name, introduction) values ('4', 'vivo', '手机品牌');


