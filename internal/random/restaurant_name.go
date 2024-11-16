package random

import "AllReady/internal/models"

func DefinitionOfRestaurantName(ds *models.DataSet, TID int) string {
	temp := ds.Tables[TID]
	rest := ds.Restaurants[temp.RID]
	return rest.Name
}
