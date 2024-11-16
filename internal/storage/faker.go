package storage

import (
	"AllReady/internal/models"
	"AllReady/internal/random"
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const num_of_nodes = 1000

func (s *PgRepository) FillAllTables(ctx context.Context) {
	FillTableUsers(ctx, s.pool)
	FillTableRestaurants(ctx, s.pool)
	FillTableReviews(ctx, s.pool)
	FillTableRTables(ctx, s.pool)
	FillTableOrders(ctx, s.pool)
	FillTableMenus(ctx, s.pool)
	FillTableDishes(ctx, s.pool)
	FillTableOrderDish(ctx, s.pool)
}

func FillTableUsers(ctx context.Context, p *pgxpool.Pool) {
	for i := 0; i < num_of_nodes; i++ {
		user := models.User{
			ID:        i,
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Telephone: gofakeit.Phone(),
			Email:     gofakeit.Email(),
			Password:  gofakeit.Password(true, false, false, false, false, 32),
			DateOfReg: gofakeit.Date(),
		}
		query := `
		INSERT INTO users (id, first_name, last_name, telephone, email, password, date_of_registration)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`
		_, err := p.Exec(ctx, query, user.ID, user.FirstName, user.LastName, user.Telephone, user.Email, user.Password, user.DateOfReg)
		if err != nil {
			log.Printf("filling table users error: %v", err)
		}
	}
}

func FillTableRestaurants(ctx context.Context, p *pgxpool.Pool) {
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
	}
}

func FillTableReviews(ctx context.Context, p *pgxpool.Pool) {
	for i := 0; i < num_of_nodes; i++ {
		review := models.Review{
			ID:         i,
			UID:        gofakeit.Number(1, num_of_nodes-1),
			RID:        gofakeit.Number(1, num_of_nodes-1),
			Rating:     gofakeit.Number(1, 5),
			Comment:    gofakeit.Paragraph(1, 3, 10, " "),
			ReviewDate: gofakeit.Date(),
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

func FillTableRTables(ctx context.Context, p *pgxpool.Pool) {

	for i := 0; i < num_of_nodes; i++ {
		rtable := models.RTable{
			ID:           i,
			RID:          gofakeit.Number(1, num_of_nodes-12),
			Seats:        gofakeit.Number(2, 10),
			Availability: gofakeit.Bool(),
			Location:     gofakeit.Word(),
		}

		query := `
		INSERT INTO restauranttables (id, restaurantid, seats, availability, location)
VALUES ($1, $2, $3, $4, $5)
`
		_, err := p.Exec(ctx, query, rtable.ID, rtable.RID, rtable.Seats, rtable.Availability, rtable.Location)
		if err != nil {
			log.Printf("filling table restauranttables error: %v", err)
		}
	}
}

func FillTableOrders(ctx context.Context, p *pgxpool.Pool) {
	for i := 0; i < num_of_nodes; i++ {
		order := models.Order{
			ID:          i,
			UID:         gofakeit.Number(1, num_of_nodes-1),                                               // Случайный ID пользователя
			RName:       gofakeit.Company(),                                                               // Случайное название ресторана
			OrderDate:   gofakeit.Date(),                                                                  // Случайная дата
			Status:      gofakeit.RandomString([]string{"pending", "confirmed", "completed", "canceled"}), // Случайный статус
			TotalAmount: gofakeit.Number(20, 500),                                                         // Сумма заказа (от 20 до 500)
			TID:         gofakeit.Number(1, num_of_nodes-1),                                               // Случайный ID стола
		}
		query := `
		INSERT INTO orders (id, userid, restaurant_name, order_date_time, order_status, total_amount, tableid)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`
		_, err := p.Exec(ctx, query, order.ID, order.UID, order.RName, order.OrderDate, order.Status, order.TotalAmount, order.TID)
		if err != nil {
			log.Printf("filling table orders error: %v", err)
		}
	}
}

func FillTableMenus(ctx context.Context, p *pgxpool.Pool) {
	for i := 0; i < num_of_nodes; i++ {
		menu := models.Menu{
			ID:   i,
			RID:  gofakeit.Number(1, num_of_nodes-1),
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
	}
}

func FillTableDishes(ctx context.Context, p *pgxpool.Pool) {
	for i := 0; i < num_of_nodes; i++ {
		dish := models.Dish{
			ID:           i,
			MID:          gofakeit.Number(1, num_of_nodes-1), // Случайный ID меню
			DishName:     gofakeit.MinecraftFood(),           // Случайное название блюда
			DishDesc:     gofakeit.Sentence(5),               // Краткое описание блюда
			Price:        gofakeit.Number(5, 100),            // Случайная цена от 5 до 100
			Availability: gofakeit.Bool(),                    // Доступность блюда
		}

		query := `
		INSERT INTO dish (id, menuid, dish_name, dish_description, price, availability)
VALUES ($1, $2, $3, $4, $5, $6)
`
		_, err := p.Exec(ctx, query, dish.ID, dish.MID, dish.DishName, dish.DishDesc, dish.Price, dish.Availability)
		if err != nil {
			log.Printf("filling table dish error: %v", err)
		}
	}
}

func FillTableOrderDish(ctx context.Context, p *pgxpool.Pool) {
	for i := 0; i < num_of_nodes; i++ {
		orderDish := models.OrderDish{
			ID:       i,
			OID:      gofakeit.Number(1, num_of_nodes-1), // Случайный ID заказа (ссылается на таблицу orders)
			DID:      gofakeit.Number(1, num_of_nodes-1), // Случайный ID блюда (ссылается на таблицу dishes)
			Quantity: gofakeit.Number(1, 5),              // Количество блюда (от 1 до 5)
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
