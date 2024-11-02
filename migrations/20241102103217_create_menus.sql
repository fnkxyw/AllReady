-- +goose Up
CREATE TABLE IF NOT EXISTS Menus (
    ID INT PRIMARY KEY,
    restaurantID INT,
    name VARCHAR(100),
    FOREIGN KEY (restaurantID) REFERENCES Restaurants(ID)
    );
-- +goose Down
DROP table Menus;