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

type Identifier interface {
    IsMatch(string) bool
    GetValue() interface{}
}

type NameIdentifier struct {
    Value string
}
type UUIDIdentifier struct {
    Value uuid.UUID
}

type NumericIdentifier struct {
    Value uint
}

func (u *UUIDIdentifier) IsMatch(s string) bool {
    val, err := uuid.FromString(s)
    if err == nil {
        u.Value = val
        return true
    }
    return false
}

func (n *NumericIdentifier) IsMatch(s string) bool {
    val, err := strconv.ParseUint(s, 10, 64)
    if err == nil {
        n.Value = uint(val)
        return true
    }
    return false
}

func (n *NameIdentifier) IsMatch(s string) bool {
    _, uidErr := uuid.FromString(s)
    _, numErr := strconv.ParseUint(s, 10, 64)
    if uidErr != nil && numErr != nil {
        n.Value = s
        return true
    }
    return false
}


func (u *UUIDIdentifier) GetValue() interface{} {
    return u.Value
}


func (n *NumericIdentifier) GetValue() interface{} {
    return n.Value
 }

func (n *NameIdentifier) GetValue() interface{} {
    return n.Value
}


func DeletePlayer(c echo.Context) error {
    // Open the database connection
    db := database.GetDB()
    id := c.Param("id")

    identifiers := []Identifier{&UUIDIdentifier{}, &NumericIdentifier{}, &NameIdentifier{}}
    
    var matchedValue interface{}
    for _, identifier := range identifiers {
        if identifier.IsMatch(id) {
            matchedValue = identifier.GetValue()
            break
        }
    }

    if matchedValue == nil {
        return c.JSON(http.StatusBadRequest, "Invalid Identifier provided")
    }

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

func handleDeleteresult(c echo.Context, result *gorm.DB) error { 
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to delete player from the database")
    }
    if result.RowsAffected == 0 {
        return c.JSON(http.StatusNotFound, "No player found in the database")
    }
    return c.JSON(http.StatusOK, "Player Deleted successfully.")
}