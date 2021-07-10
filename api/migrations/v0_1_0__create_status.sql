-- +migrate Up
create table `status` (
    `id` int(11) not null auto_increment,
    `title` varchar(256) not null unique,
    primary key (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `status`;
