-- +goose Up
CREATE TABLE IF NOT EXISTS RestaurantTables (
    ID INT PRIMARY KEY,
    restaurantID INT,
    seats INT,
    availability BOOLEAN,
    location VARCHAR(50),
    FOREIGN KEY (restaurantID) REFERENCES restaurants(ID)
);

-- +goose Down
DROP TABLE RestaurantTables;