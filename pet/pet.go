package pet

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/lindsay0416/server/database"
)

// Object
type Pet struct {
	gorm.Model
	ID     int    `json:"id"`
	Breed  string `json:"breed"`
	Name   string `json:"name,omitempty"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// GET http://localhost:8080/api/pet/:id?name="Zhu"
// Handler / Controller
func GetPetURL(c echo.Context) error {
	// Pet ID from path `pet/:id`
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	name := c.QueryParam("name")
	if name == "" {
		return errors.New("empty name query param")
	}

	breed := ""

	switch id {
	case 1:
		breed = "Ragdoll" + name
	case 2:
		breed = "Tabby " + name
	case 3:
		breed = "Short Hair" + name
	}

	resp := Pet{
		Breed: breed,
		ID:    id,
		Name:  name,
	}

	return c.JSON(http.StatusOK, resp)

}

// GetPets func
// Get all records of pets in database
func GetPets(c echo.Context) error {
	db := database.DBConn
	var pets []Pet
	db.Find(pets)
	return c.JSON(http.StatusOK, pets)
}

// GetPet func
// GetPet by id
func GetPet(c echo.Context) error {

	id := c.Param("id")
	db := database.DBConn
	var pet Pet
	db.Find(pet, id)
	return c.JSON(http.StatusOK, pet)
}

// POST http://localhost:8080/api/pet
func NewPet(c echo.Context) error {
	db := database.DBConn
	var pet Pet
	pet.Name = "Happy"
	pet.Age = 1
	pet.Breed = "Domestic Short Hair"
	db.Create(pet)
	return c.JSON(http.StatusOK, pet)

}

func DeletePet(c echo.Context) error {
	idStr := c.Param("id")
	db := database.DBConn
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	var pet Pet
	db.First(pet, id)
	if pet.Name == "" {
		return errors.New("empty name query param")
	}
	db.Delete(pet)
	return c.JSON(http.StatusOK, "Pet successfully deleted")
}
