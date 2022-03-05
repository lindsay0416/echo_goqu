package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lindsay0416/server/database"
	"github.com/lindsay0416/server/pet"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "pets.db")
	if err != nil {
		panic("Failed to connect to database.")
	}
	fmt.Println("Database connection success.")
}

// Setup Routes
// http://localhost:8080/api/pet
func setupRoutes(e *echo.Echo) {
	e.GET("api/pet", pet.GetPets)
	e.GET("api/pet/:id", pet.GetPet)
	e.GET("api/pet/:id", pet.GetPetURL)
	e.POST("api/pet", pet.NewPet)
	e.DELETE("api/pet/:id", pet.DeletePet)
}

func main() {
	// Echo instance
	e := echo.New()
	// initDatabase()
	setupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
	// defer database.DBController.DBConn.Close()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

}
