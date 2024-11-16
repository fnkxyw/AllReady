package models

import "time"

//all models for our db

type User struct {
	ID        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Telephone string    `db:"telephone"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	DateOfReg time.Time `db:"date_of_registration"`
}

type Review struct {
	ID         int       `db:"id"`
	UID        int       `db:"userid"'`
	RID        int       `db:"restaurantid"`
	Rating     int       `db:"rating"`
	Comment    string    `db:"comment"`
	ReviewDate time.Time `db:"review_date"`
}

type RTable struct {
	ID           int    `db:"id"`
	RID          int    `db:"restaurantid"`
	Seats        int    `db:"seats"`
	Availability bool   `db:"availability"`
	Location     string `db:"location"`
}

type Restaurant struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	Telephone   string `db:"telephone"`
	Rating      int    `db:"rating"`
	Description string `db:"description"`
	WorkTime    string `db:"work_time"`
}

type Order struct {
	ID          int       `db:"id"`
	UID         int       `db:"userid"`
	RName       string    `db:"restaurant_name"`
	OrderDate   time.Time `db:"order_date_time"`
	Status      string    `db:"order_status"`
	TotalAmount int       `db:"total_amount"`
	TID         int       `db:"tableid"`
}

type OrderDish struct {
	ID       int `db:"id"`
	OID      int `db:"orderid"`
	DID      int `db:"dishid"`
	Quantity int `db:"quantity"`
}

type Menu struct {
	ID   int    `db:"id"`
	RID  int    `db:"rastaurantid"`
	Name string `db:"name"`
}

type Dish struct {
	ID           int    `db:"id"`
	MID          int    `db:"menuid"`
	DishName     string `db:"dish_name"`
	DishDesc     string `db:"dish_description"`
	Price        int    `db:"price"`
	Availability bool   `db:"availability"`
}
