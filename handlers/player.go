package handlers

import (
	"encoding/json"
	"log"

	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Player(player models.Player) bool {

	log.Printf("We are in Player Handler")
	prettyPlayerArg, _ := PrettyStruct(player)
	log.Printf("Player Arg  %v", prettyPlayerArg)

	db := database.Open()

	// Initialize empty database player model to bind into from db query
	databasePlayerModel := &models.Player{}

	var result *gorm.DB

	if viper.GetBool("DEBUG") {
		log.Printf("Querying player by Name because we are in DEBUG mode")

		// Get the player by name in DEBUG mode for easier debugging and avoid the Userservice dependency
		result = db.Preload(clause.Associations).Model(&models.Player{}).Where(&models.Player{Name: player.Name}).First(&databasePlayerModel)
	} else {
		log.Printf("Querying database player by UserID")
		// Query db with the UserID from the passed in player model to find correct player
		result = db.Preload(clause.Associations).Model(&models.Player{}).Where(&models.Player{UserID: player.UserID}).First(&databasePlayerModel)
	}

	if result.Error != nil {
		// TODO: handle error
		panic(result.Error)
	}

	prettyDatabasePlayerModel, _ := PrettyStruct(databasePlayerModel)
	prettyPosition, _ := PrettyStruct(databasePlayerModel.Position)
	log.Printf("Query result for databasePlayerModel %v", prettyDatabasePlayerModel)
	log.Printf("Query result for Position %v", prettyPosition)

	log.Printf("Updating the databasePlayerModel with the player from the function arg")

	databasePlayerModel.Position = player.Position
	databasePlayerModel.Position.ID = uint(databasePlayerModel.PositionID)

	log.Printf("Updated the databasePlayerModel the Data should look differently now!")

	prettyPositionUpdated, _ := PrettyStruct(databasePlayerModel.Position)
	log.Printf("Updated Position %v", prettyPositionUpdated)

	// // Start Association Mode
	// var playerAssociation models.Player
	// assocField := "Position"
	// db.Model(&playerAssociation).Association(assocField)
	// // `user` is the source model, it must contains primary key
	// // `Languages` is a relationship's field name
	// // If the above two requirements matched, the AssociationMode should be started successfully, or it should return error
	// err := db.Model(&playerAssociation).Association(assocField).Error
	// if err != nil {
	// 	log.Printf("Association error: %v", err)
	// }

	// positions := &models.Position{}
	// db.Model(databasePlayerModel).Association("Position").Find(&positions)

	// prettyASOCPosition, _ := PrettyStruct(positions)

	// log.Printf("Associated positions: %v", prettyASOCPosition)
	// log.Printf("Setting New Data for Player")
	// log.Printf("Setting New Data for Player")
	// log.Printf("Setting New Data for Player")
	// log.Printf("databasePlayerModel.PositionID:%v", uint(databasePlayerModel.PositionID))

	// player.Position.ID = uint(databasePlayerModel.PositionID)

	// log.Printf("player.Position.ID:%v", player.Position.ID)
	// log.Printf("player.Position:%v", player.Position)

	// // TODO: Check Updating Functionality and behaviour(also in regards to partial updates)
	// // Updating the database player model with the new data from player argument
	// databasePlayerModel.Position.Vector3 = player.Position.Vector3
	// databasePlayerModel.Position.Vector3.X = player.Position.Vector3.X

	// // databasePlayerModel.Rotation = player.Rotation
	// // databasePlayerModel.Scale = player.Scale

	log.Printf("Updated database player model Vector3X: %v", databasePlayerModel.Position.Vector3.X)

	log.Printf("Saving database player")
	// db.Save(databasePlayerModel)
	// db.Save(databasePlayerModel.Position)
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&databasePlayerModel)
	log.Printf("Returning ")

	return true

}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
