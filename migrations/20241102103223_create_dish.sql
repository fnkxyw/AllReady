-- +goose Up
CREATE TABLE Dish (
    ID INT PRIMARY KEY,
    menuID INT,
    dish_name VARCHAR(100),
    dish_description TEXT,
    price DECIMAL(10, 2),
    availability BOOLEAN,
    FOREIGN KEY (menuID) REFERENCES Menus(ID)
);

-- +goose Down
DROP TABLE Dish;