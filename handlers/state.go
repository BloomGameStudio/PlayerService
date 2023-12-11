package handlers

import (
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func State(state models.State, c echo.Context) error {
	// c echo.Context is only used to obtain the logger to provide unified logging
	logger := c.Logger()

	logger.Debug("We are in State Handler")
	logger.Debugf("state Arg: %+v", state)

	db := database.GetDB()

	// Initialize empty database state model to bind into from db query
	databaseStateModel := &models.State{}

	var result *gorm.DB

	logger.Debug("Querying state from the database by ID")
	// Query db with the ID from the passed in state model to find correct player
	queryState := &models.State{}
	queryState.ID = state.ID

	result = db.Model(&models.State{}).Where(queryState).First(databaseStateModel)

	if result.Error != nil {
		logger.Error(result.Error)
		return result.Error
	}

	logger.Debugf("Query result for databaseStateModel: %+v", databaseStateModel)
	logger.Debug("Updating the existing databaseStateModel")

	// Update the databaseStateModel with the provided state values
	databaseStateModel.StateID = state.StateID
	databaseStateModel.Value = state.Value

	logger.Debugf("Updated databaseStateModel: %+v", databaseStateModel)

	logger.Debug("Saving database State")
	result = db.Save(databaseStateModel)

	if result.Error != nil {
		logger.Error(result.Error)
		return result.Error
	}

	return nil
}