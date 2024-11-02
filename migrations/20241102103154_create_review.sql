-- +goose Up
CREATE TABLE IF NOT EXISTS Review (
    ID INT PRIMARY KEY,
    userID INT,
    restaurantID INT,
    rating INT CHECK (Rating BETWEEN 1 AND 5),
    comment TEXT,
    review_date DATE,
    FOREIGN KEY (userID) REFERENCES Users(ID),
    FOREIGN KEY (RestaurantID) REFERENCES Restaurants(ID)
    );

-- +goose Down
DROP table Review;