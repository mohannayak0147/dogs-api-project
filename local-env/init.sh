#!/bin/sh

mysql -h localhost -u root -P 3306 -phello987 <<EOF
drop database if exists demo01;
create database demo01;
use demo01;

drop table if exists dogs;
create table dogs (
    id int(11) AUTO_INCREMENT PRIMARY KEY,
    breed          varchar(256)      NOT NULL UNIQUE,
    data           varchar(256)      NOT NULL,
    updated_at    datetime NOT NULL,
    created_at      datetime NOT NULL
)
ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

EOF
