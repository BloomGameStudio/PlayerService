package controllers

import (
	"net/http"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/BloomGameStudio/PlayerService/publicModels"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

func CreatePlayer(c echo.Context) error {

	// Creates a new Player in the Service
	// Expects a publicModel or a model dot Player object in the body

	c.Logger().Debug("In CreatePlayer")

	// Initializer request player to bind into
	reqPlayer := publicModels.Player{}

	err := c.Bind(&reqPlayer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	c.Logger().Debugf("Bound the player struct: %v", reqPlayer)

	if !reqPlayer.IsValid() {
		c.Logger().Debug("reqPlayer is NOT valid")
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	c.Logger().Debug("reqPlayer is valid")

	c.Logger().Debug("Initializing and populating player model!")
	// Use dot annotation for promoted aka embedded fields.
	playerModel := &models.Player{}

	// TODO: Set UserID
	// Note: TMP HOTFIX To create users in staging env without having to enter server or container
	if viper.GetBool("DEBUG") {
		// Sleep for 1 second in DEBUG mode to not get fludded with data
		playerModel.UserID = uuid.NewV4()
	}

	playerModel.Name = reqPlayer.Name
	playerModel.Position.Position = reqPlayer.Position
	playerModel.Rotation.Rotation = reqPlayer.Rotation
	playerModel.Scale.Scale = reqPlayer.Scale

	if !playerModel.IsValid() {
		c.Logger().Debug("playerModel is NOT valid")
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	c.Logger().Debug("playerModel is valid")

	// Save to db
	db := database.GetDB()
	db.Create(playerModel)

	c.Logger().Debug("playerModel is saved. Returning")

	return c.JSON(http.StatusCreated, playerModel)
}

//Define a response struct for control over what gets serialized?

func GetPlayer(c echo.Context) error {
	// Open the database connection
	db := database.GetDB()

	// Read the "active" query parameter from the URL
	activeParam := c.QueryParam("active")

	// Initialize a variable to store the filter value
	var active bool

	// Check if the "active" query parameter is provided and parse it as a boolean
	switch activeParam {
	case "true":
		active = true
	case "false":
		active = false
	case "":
		//handle some kind of things

	default:
		return c.JSON(http.StatusBadRequest, "Invalid 'active' parameter value. Use 'true' or 'false'.")
	}

	// Build the query based on the "active" filter
	queryPlayer := &models.Player{}
	queryPlayer.Active = active
	players := &[]models.Player{}
	if err := db.Preload(clause.Associations).Where(queryPlayer.Active).Find(players).Error; err != nil {
		c.Logger().Error("Failed to retrieve players from the database")
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve players from the database")
	}
	// Return the list of players as a JSON response
	return c.JSON(http.StatusOK, players)
}
