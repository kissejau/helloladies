CREATE TABLE cities (
    id VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    confirmed BOOLEAN
);

CREATE TABLE univs (
    id VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    city_id VARCHAR REFERENCES cities (id) ON DELETE NO ACTION,
    confirmed BOOLEAN
);

CREATE TABLE teachers (
    id VARCHAR PRIMARY KEY,
    code VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    patronymic VARCHAR NOT NULL,
    degree VARCHAR,
    univ_id VARCHAR REFERENCES univs (id) ON DELETE NO ACTION,
    confirmed BOOLEAN
);

CREATE TABLE reviews (
    id VARCHAR PRIMARY KEY,
    description VARCHAR NOT NULL,
    rating INTEGER NOT NULL,
    teacher_id VARCHAR REFERENCES teachers (id) ON DELETE NO ACTION,
    user_id VARCHAR REFERENCES users (id) ON DELETE NO ACTION
);



