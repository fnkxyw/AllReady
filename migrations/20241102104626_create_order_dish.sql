-- +goose Up
CREATE TABLE Order_Dish (
    ID INT PRIMARY KEY,
    orderID INT,
    dishID INT,
    quantity INT,
    FOREIGN KEY (orderID) REFERENCES Orders(ID),
    FOREIGN KEY (dishID) REFERENCES Dish(ID)
);

-- +goose Down
DROP TABLE Order_Dish;