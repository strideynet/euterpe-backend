CREATE TABLE teapots (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    radius FLOAT NOT NULL,
    height FLOAT NOT NULL,
    create_time TIMESTAMP NOT NULL,
    update_time TIMESTAMP NOT NULL
);