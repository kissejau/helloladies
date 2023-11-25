CREATE TABLE IF NOT EXISTS users(
    id VARCHAR PRIMARY KEY,
    email VARCHAR(50),
    password VARCHAR,
    name VARCHAR(40),
    birth_date timestamp
)