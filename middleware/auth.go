package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/suradidchao/listenfield/handler"
)

// AuthorizeFarmAccess is the middleware function for authorization.
func AuthorizeFarmAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		farmID, err := strconv.Atoi(c.Param("farm_id"))
		if err != nil {
			fmt.Printf("Err: %s", err)
			return c.JSON(http.StatusInternalServerError, handler.Response{Message: "Invalid farm id"})
		}

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		ownedFarmIDs := []int{}
		ownerFarmIDsInterface := claims["ownedFarmIds"].([]interface{})
		for _, ownedFarmID := range ownerFarmIDsInterface {
			ownedFarmIDs = append(ownedFarmIDs, int(ownedFarmID.(float64)))
		}

		accessible := canAccessFarm(farmID, ownedFarmIDs)
		if accessible {
			if err := next(c); err != nil {
				c.Error(err)
			}
		}
		return c.JSON(http.StatusInternalServerError, handler.Response{Message: "Not Authorize"})
	}
}

func canAccessFarm(farmID int, farmIDs []int) bool {
	for _, fid := range farmIDs {
		if farmID == fid {
			return true
		}
	}
	return false
}
