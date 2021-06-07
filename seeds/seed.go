package seeds

import (
	"github/sing3demons/go-fiber-admin/database"
	"github/sing3demons/go-fiber-admin/models"
	"strconv"

	"github.com/bxcodec/faker/v3"
	"github.com/ian-kent/go-log/log"
)

func CreatePermisstion() {
	perm := [...]models.Permission{
		{Name: "view_users"},
		{Name: "edit_users"},
		{Name: "view_roles"},
		{Name: "edit_roles"},
		{Name: "view_products"},
		{Name: "edit_products"},
		{Name: "view_orders"},
		{Name: "edit_orders"},
	}
	database.DB.Migrator().DropTable(&models.Permission{})
	database.DB.AutoMigrate(models.Permission{})
	database.DB.Create(&perm)

}

func Load() {

	database.DB.Migrator().DropTable(&models.Product{}, &models.Order{})
	database.DB.AutoMigrate(models.Product{}, &models.Order{})
	log.Info("Creating products...")
	numOfProduct := 50
	// products := make([]models.Product, 0, numOfProduct)
	for i := 1; i <= numOfProduct; i++ {
		product := models.Product{
			Title:       faker.Word(),
			Description: faker.Paragraph(),
			Price:       float64(faker.UnixTime()),
			Image:       "https://source.unsplash.com/random/300x200?" + strconv.Itoa(i),
		}
		database.DB.Create(&product)
		// products = append(products, product)
	}

	log.Info("Creating order...")
	numOFOrder := 10
	for i := 1; i <= numOFOrder; i++ {
		order := models.Order{
			Email:     faker.Email(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Total:     1,
		}
		database.DB.Create(&order)
	}

}
