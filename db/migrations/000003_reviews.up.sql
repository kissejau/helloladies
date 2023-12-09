CREATE TABLE cities (
    id VARCHAR PRIMARY KEY,
    code VARCHAR NOT NULL UNIQUE,
    title VARCHAR NOT NULL UNIQUE,
    confirmed BOOLEAN
);

CREATE TABLE univs (
    id VARCHAR PRIMARY KEY,
    code VARCHAR NOT NULL UNIQUE,
    title VARCHAR NOT NULL,
    city_id VARCHAR REFERENCES cities (id) ON DELETE NO ACTION,
    confirmed BOOLEAN
);

CREATE TABLE teachers (
    id VARCHAR PRIMARY KEY,
    code VARCHAR NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    patronymic VARCHAR NOT NULL,
    degree VARCHAR,
    univ_id VARCHAR REFERENCES univs (id) ON DELETE NO ACTION,
    confirmed BOOLEAN
);

CREATE TABLE reviews (
    id VARCHAR PRIMARY KEY,
    code VARCHAR NOT NULL UNIQUE,
    description VARCHAR NOT NULL,
    rating INTEGER NOT NULL,
    teacher_id VARCHAR REFERENCES teachers (id) ON DELETE NO ACTION,
    user_id VARCHAR REFERENCES users (id) ON DELETE NO ACTION
);



