CREATE KEYSPACE IF NOT EXISTS post_service WITH replication = {
    'class': 'SimpleStrategy', 
    'replication_factor': 1
};

USE post_service;

CREATE TYPE IF NOT EXISTS historical_interests (
    id text,
    interest text
);

CREATE TABLE IF NOT EXISTS posts (
    id text PRIMARY KEY,
    username text,
    first_name text,
    last_name text,
    title text,
    subtitle text,
    abstract text,
    post_body text,
    validity_rating double,
    historical_interests map<text, frozen<historical_interests>>
);