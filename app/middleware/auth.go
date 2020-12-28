package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/suradidchao/listenfield/app/handler"
)

// AuthorizeFarmOwnerAccess is the middleware function for authorization.
func AuthorizeFarmOwnerAccess(next echo.HandlerFunc) echo.HandlerFunc {
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

		if !accessible {
			return c.JSON(http.StatusUnauthorized, handler.Response{Message: "Unauthorized"})
		}

		err = next(c)
		if err != nil {
			c.Error(err)
		}
		return nil
	}
}

// AuthorizeFarmOwnerAndWorkerAccess is the middleware function for authorization.
func AuthorizeFarmOwnerAndWorkerAccess(next echo.HandlerFunc) echo.HandlerFunc {
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

		workingFarmIDs := []int{}
		workingFarmIDsInterface := claims["workingFarmIds"].([]interface{})
		for _, workingFarmID := range workingFarmIDsInterface {
			workingFarmIDs = append(workingFarmIDs, int(workingFarmID.(float64)))
		}

		accessible := canAccessFarm(farmID, ownedFarmIDs) || canAccessFarm(farmID, workingFarmIDs)

		if !accessible {
			return c.JSON(http.StatusUnauthorized, handler.Response{Message: "Unauthorized"})
		}

		err = next(c)
		if err != nil {
			c.Error(err)
		}
		return nil
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
