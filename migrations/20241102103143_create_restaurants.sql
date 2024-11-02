-- +goose Up
CREATE TABLE IF NOT EXISTS Restaurants (
    ID INT PRIMARY KEY,
    name VARCHAR(100),
    address VARCHAR(200),
    telephone VARCHAR(20),
    rating DECIMAL(2, 1),
    description TEXT,
    work_time VARCHAR(100)
    );
-- +goose Down
DROP table Restaurants;