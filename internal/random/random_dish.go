package random

import (
	"AllReady/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"math/rand"
)

//func for correct dish pick

func RandomDishIDForOrderDishTable(ctx context.Context, ds *models.DataSet, OID int, pr *pgxpool.Pool) int {
	temp := ds.Tables[ds.Orders[OID].TID]
	rest := ds.Restaurants[temp.RID]

	rows, err := pr.Query(ctx, `SELECT * FROM menus WHERE restaurantid=$1`, rest.ID)
	if err != nil {
		log.Printf("GetMenusWithRestName error: %v", err)
	}

	menus, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Menu])
	if err != nil {
		log.Printf("CollectRows error: %v", err)
	}
	randomIndex := rand.Intn(len(menus))
	RMenu := menus[randomIndex]

	rows, err = pr.Query(ctx, `SELECT * FROM dish WHERE menuid=$1`, RMenu.ID)
	if err != nil {
		log.Println("GetMenusWithRestName error")
	}

	dishes, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Dish])
	if err != nil {
		log.Println("CollectRows error")

	}

	randomIndex = rand.Intn(len(dishes))
	RDish := dishes[randomIndex]
	return RDish.ID
}
