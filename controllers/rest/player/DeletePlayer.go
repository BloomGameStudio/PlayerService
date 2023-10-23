package player

import (
	"net/http"
	"github.com/BloomGameStudio/PlayerService/database"
	"github.com/BloomGameStudio/PlayerService/models"
	"github.com/labstack/echo/v4"
    uuid "github.com/satori/go.uuid"
	"strconv"
    "gorm.io/gorm"
)

//Interface to validate and fetch values
type Identifier interface {
    IsMatch(string) bool
    GetValue() interface{}
}

//Used to determine if the input is a players name
type NameIdentifier struct {
    Value string
}

//Used to determine if the input is a player UUID
type UUIDIdentifier struct {
    Value uuid.UUID
}

//Used to determine if the input is a players ID
type NumericIdentifier struct {
    Value uint
}

//Checks if the given string can be parsed as a UUID
func (u *UUIDIdentifier) IsMatch(s string) bool {
    val, err := uuid.FromString(s)
    if err == nil {
        u.Value = val
        return true
    }
    return false
}

//Checks if the given string can be parsed as an uint
func (n *NumericIdentifier) IsMatch(s string) bool {
    val, err := strconv.ParseUint(s, 10, 64)
    if err == nil {
        n.Value = uint(val)
        return true
    }
    return false
}

//Checks if the given string is neither a UUID nor an uint
func (n *NameIdentifier) IsMatch(s string) bool {
    _, uidErr := uuid.FromString(s)
    _, numErr := strconv.ParseUint(s, 10, 64)
    if uidErr != nil && numErr != nil {
        n.Value = s
        return true
    }
    return false
}

//Returns the parsed UUID value
func (u *UUIDIdentifier) GetValue() interface{} {
    return u.Value
}

//Returns the parsed numeric value
func (n *NumericIdentifier) GetValue() interface{} {
    return n.Value
 }

 //Returns the name value
func (n *NameIdentifier) GetValue() interface{} {
    return n.Value
}


//Delete the player
func DeletePlayer(c echo.Context) error {
    // Open the database connection
    db := database.GetDB()
    id := c.Param("id")

    //Create a list of identifiers
    identifiers := []Identifier{&UUIDIdentifier{}, &NumericIdentifier{}, &NameIdentifier{}}
    
    //Loop through to determine type of identifier
    var matchedValue interface{}
    for _, identifier := range identifiers {
        if identifier.IsMatch(id) {
            matchedValue = identifier.GetValue()
            break
        }
    }

    //No match error
    if matchedValue == nil {
        return c.JSON(http.StatusBadRequest, "Invalid Identifier provided")
    }

    //use the matched value to delete the appropriate player entry
    switch v := matchedValue.(type){
    case uuid.UUID:
        result := db.Where("user_id = ?", v).Delete(&models.Player{})
        return handleDeleteresult(c, result)
    case uint:
        result := db.Where("id = ?", v).Delete(&models.Player{})
        return handleDeleteresult(c, result)
    case string:
        result := db.Where("name = ?", v).Delete(&models.Player{})
        return handleDeleteresult(c, result)
    default:
        return c.JSON(http.StatusBadRequest, "Invalid Identifier")
    }
}

//Take result of database delete operation and returns a HTTP response
func handleDeleteresult(c echo.Context, result *gorm.DB) error { 
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to delete player from the database")
    }
    if result.RowsAffected == 0 {
        return c.JSON(http.StatusNotFound, "No player found in the database")
    }
    return c.JSON(http.StatusOK, "Player Deleted successfully.")
}