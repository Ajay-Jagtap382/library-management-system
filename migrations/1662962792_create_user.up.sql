CREATE TABLE user
(
    `id` varchar(10) NOT NULL,
    `first_name` varchar(15) NOT NULL,
    `last_name` varchar(15) NOT NULL,
    `mobile_num` varchar(10) NOT NULL ,
    `email` varchar(20) NOT NULL,
    `password` varchar(15) NOT NULL,
    `gender` varchar(5) NOT NULL,
    PRIMARY KEY(`id`),
    UNIQUE(`email`)

);