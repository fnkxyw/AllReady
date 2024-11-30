package storage

import (
	"AllReady/internal/models"
	"AllReady/internal/random"
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

//file with functions for filling tables / gofakeit

const num_of_nodes = 10000

//method for filling all tables in db

func (s *PgRepository) FillAllTables(ctx context.Context) {
	//simple methods for every table
	ds := models.DataSet{}

	FillTableUsers(ctx, s.pool, &ds)
	FillTableRestaurants(ctx, s.pool, &ds)
	FillTableReviews(ctx, s.pool, &ds)
	FillTableRTables(ctx, s.pool, &ds)
	FillTableMenus(ctx, s.pool, &ds)
	FillTableDishes(ctx, s.pool, &ds)
	FillTableOrders(ctx, s.pool, &ds)
	FillTableOrderDish(ctx, s.pool, &ds)
}

func FillTableUsers(ctx context.Context, p *pgxpool.Pool, ds *models.DataSet) {
	for i := 0; i < num_of_nodes; i++ {
		user := models.User{
			ID:        i,
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Telephone: gofakeit.Phone(),
			Email:     gofakeit.Email(),
			Password:  gofakeit.Password(true, false, false, false, false, 32),
			DateOfReg: random.GenerateDateAfter2010(),
		}
		query := `
		INSERT INTO users (id, first_name, last_name, telephone, email, password, date_of_registration)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`
		_, err := p.Exec(ctx, query, user.ID, user.FirstName, user.LastName, user.Telephone, user.Email, user.Password, user.DateOfReg)
		if err != nil {
			log.Printf("filling table users error: %v", err)
		}
		ds.DatesOfRegistration = append(ds.DatesOfRegistration, user.DateOfReg)

	}
}

func FillTableRestaurants(ctx context.Context, p *pgxpool.Pool, ds *models.DataSet) {
	for i := 0; i < num_of_nodes; i++ {
		restaurant := models.Restaurant{
			ID:          i,
			Name:        gofakeit.Company(),
			Address:     gofakeit.Address().Address,
			Telephone:   gofakeit.Phone(),
			Rating:      gofakeit.Number(1, 5),
			Description: gofakeit.Sentence(10),
			WorkTime:    random.GenerateWorkHours(),
		}
		query := `
		INSERT INTO restaurants (id, name, address, telephone, rating, description, work_time)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`
		_, err := p.Exec(ctx, query, restaurant.ID, restaurant.Name, restaurant.Address, restaurant.Telephone, restaurant.Rating, restaurant.Description, restaurant.WorkTime)
		if err != nil {
			log.Printf("filling table restaurants error: %v", err)
		}
		ds.Restaurants = append(ds.Restaurants, restaurant)

	}
}

func FillTableReviews(ctx context.Context, p *pgxpool.Pool, ds *models.DataSet) {
	times := random.GenerateRandomReviewDate(ds.DatesOfRegistration)
	for i := 0; i < num_of_nodes; i++ {
		UID := gofakeit.Number(1, num_of_nodes-1)
		review := models.Review{
			ID:         i,
			UID:        UID,
			RID:        gofakeit.Number(1, num_of_nodes-1),
			Rating:     gofakeit.Number(1, 5),
			Comment:    gofakeit.Paragraph(1, 3, 10, " "),
			ReviewDate: times[UID],
		}
		query := `
		INSERT INTO review (id, userid, restaurantid, rating, comment, review_date)
VALUES ($1, $2, $3, $4, $5, $6)
`
		_, err := p.Exec(ctx, query, review.ID, review.UID, review.RID, review.Rating, review.Comment, review.ReviewDate)
		if err != nil {
			log.Printf("filling table reviews error: %v", err)
		}
	}
}

func FillTableRTables(ctx context.Context, p *pgxpool.Pool, ds *models.DataSet) {

	for i := 0; i < num_of_nodes*5; i++ {
		rtable := models.RTable{
			ID:           i,
			RID:          gofakeit.Number(1, num_of_nodes-1),
			Seats:        gofakeit.Number(2, 10),
			Availability: gofakeit.Bool(),
			Location:     random.GenerateRandomLocation(),
		}

		query := `
		INSERT INTO restauranttables (id, restaurantid, seats, availability, location)
VALUES ($1, $2, $3, $4, $5)
`
		_, err := p.Exec(ctx, query, rtable.ID, rtable.RID, rtable.Seats, rtable.Availability, rtable.Location)
		if err != nil {
			log.Printf("filling table restauranttables error: %v", err)
		}
		ds.Tables = append(ds.Tables, rtable)
	}
}

func FillTableMenus(ctx context.Context, p *pgxpool.Pool, ds *models.DataSet) {
	count := 0
	for _, restaurant := range ds.Restaurants {
		for i := 0; i < 5; i++ {
			menu := models.Menu{
				ID:   count,
				RID:  restaurant.ID,
				Name: gofakeit.Company(),
			}
			query := `
			INSERT INTO menus (id, restaurantid, name)
			VALUES ($1, $2, $3)
			`
			_, err := p.Exec(ctx, query, menu.ID, menu.RID, menu.Name)
			if err != nil {
				log.Printf("filling table menu error: %v", err)
			}

			ds.Menus = append(ds.Menus, menu)
			count++
		}
	}
}

func FillTableOrders(ctx context.Context, p *pgxpool.Pool, ds *models.DataSet) {
	times := random.GenerateRandomOrderDate(ds.DatesOfRegistration)
	for i := 0; i < num_of_nodes; i++ {
		UID := gofakeit.Number(1, num_of_nodes-1)
		TID := gofakeit.Number(1, num_of_nodes-1)
		order := models.Order{
			ID:          i,
			UID:         UID,
			RName:       random.DefinitionOfRestaurantName(ds, TID),
			OrderDate:   times[UID],
			Status:      gofakeit.RandomString([]string{"pending", "confirmed", "completed", "canceled"}),
			TotalAmount: gofakeit.Number(20, 500),
			TID:         TID,
		}
		query := `
		INSERT INTO orders (id, userid, restaurant_name, order_date_time, order_status, total_amount, tableid)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`
		_, err := p.Exec(ctx, query, order.ID, order.UID, order.RName, order.OrderDate, order.Status, order.TotalAmount, order.TID)
		if err != nil {
			log.Printf("filling table orders error: %v", err)
		}
		ds.Orders = append(ds.Orders, order)
	}
}

func FillTableDishes(ctx context.Context, p *pgxpool.Pool, ds *models.DataSet) {
	count := 0
	for _, menu := range ds.Menus {
		for i := 0; i < 5; i++ {
			dish := models.Dish{
				ID:           count,
				MID:          menu.ID,
				DishName:     gofakeit.MinecraftFood(),
				DishDesc:     gofakeit.Sentence(5),
				Price:        gofakeit.Number(5, 100),
				Availability: gofakeit.Bool(),
			}

			query := `
			INSERT INTO dish (id, menuid, dish_name, dish_description, price, availability)
			VALUES ($1, $2, $3, $4, $5, $6)
			`
			_, err := p.Exec(ctx, query, dish.ID, dish.MID, dish.DishName, dish.DishDesc, dish.Price, dish.Availability)
			if err != nil {
				log.Printf("filling table dish error: %v", err)
			}
			count++
		}
	}
}

func FillTableOrderDish(ctx context.Context, p *pgxpool.Pool, ds *models.DataSet) {
	for i := 0; i < num_of_nodes; i++ {
		OID := gofakeit.Number(1, num_of_nodes-1)
		DID := random.RandomDishIDForOrderDishTable(ctx, ds, OID, p)
		orderDish := models.OrderDish{
			ID:       i,
			OID:      OID,
			DID:      DID,
			Quantity: gofakeit.Number(1, 5),
		}

		query := `
		INSERT INTO order_dish (id, orderid, dishid, quantity)
VALUES ($1, $2, $3, $4)
`
		_, err := p.Exec(ctx, query, orderDish.ID, orderDish.OID, orderDish.DID, orderDish.Quantity)
		if err != nil {
			log.Printf("filling table order_dishes error: %v", err)
		}
	}
}
