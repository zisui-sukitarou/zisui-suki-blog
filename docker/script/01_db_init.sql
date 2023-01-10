/*
CREATE DATABASE IF NOT EXISTS blog_db;
USE blog_db;

DROP TABLE IF EXISTS blogs;
CREATE TABLE test(
    blog_id char(26) NOT NULL,
    user_id char(26) NOT NULL,
    content text NOT NULL,
    title text NOT NULL,
    abstract text NOT NULL,
    evaluation int unsigned NOT NULL,
    status int unsigned NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (blog_id)
);

DROP TABLE IF EXISTS drafts;
CREATE TABLE test(
    draft_id char(26) NOT NULL,
    user_id char(26) NOT NULL,
    content text NOT NULL,
    title text NOT NULL,
    abstract text NOT NULL,
    evaluation int unsigned NOT NULL,
    status int unsigned NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (draft_id)
);

DROP TABLE IF EXISTS tags;
CREATE TABLE test(
    tag_name varchar(32) NOT NULL,
    icon text,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
);


CREATE DATABASE IF NOT EXISTS id_db;
USE id_db;

DROP TABLE IF EXISTS test;
CREATE TABLE test(
    user_id char(26) NOT NULL,
    name varchar(15) UNIQUE NOT NULL,
    display_name varchar(32) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(60) NOT NULL,
    icon text NOT NULL,
    status int unsigned NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    PRIMARY KEY (user_id)
);*/