-- +goose Up
CREATE TABLE IF NOT EXISTS Orders (
    ID INT PRIMARY KEY,
    userID INT,
    restaurant_name VARCHAR(100),
    order_date_time TIMESTAMP,
    order_status VARCHAR(50),
    total_amount DECIMAL(10, 2),
    tableID INT,
    FOREIGN KEY (userID) REFERENCES Users(ID),
    FOREIGN KEY (tableID) REFERENCES RestaurantTables(ID)
    );

-- +goose Down
DROP TABLE Orders;