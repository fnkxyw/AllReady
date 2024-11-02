-- +goose Up
CREATE TABLE  IF NOT EXISTS Users(
    ID INT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    telephone VARCHAR(20),
    email VARCHAR(100),
    password VARCHAR(100),
    date_of_registration DATE
    );

-- +goose Down
DROP TABLE Users;