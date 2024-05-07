package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xops-infra/jms/app"
	"github.com/xops-infra/jms/core/db"
	"github.com/xops-infra/noop/log"
)

// @Summary List profile
// @Description List profile
// @Tags profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []db.Profile
// @Router /api/v1/profile [get]
func listProfile(c *gin.Context) {
	profiles, err := app.App.DBService.ListProfile()
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, profiles)
}

// @Summary Create profile
// @Description Create profile
// @Tags profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param profile body db.CreateProfileRequest true "profile"
// @Success 200 {string} string
// @Router /api/v1/profile [post]
func createProfile(c *gin.Context) {
	var req db.CreateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	id, err := app.App.DBService.CreateProfile(req)
	if err != nil {
		log.Errorf("create profile error: %v", err)
		c.JSON(500, err.Error())
		return
	}
	c.String(200, id)
}

// @Summary Update profile
// @Description Update profile
// @Tags profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param profile body db.CreateProfileRequest true "profile"
// @Param uuid path string true "profile uuid"
// @Success 200 {string} string
// @Router /api/v1/profile/:uuid [put]
func updateProfile(c *gin.Context) {
	var req db.CreateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(400, "uuid is required")
		return
	}
	err := app.App.DBService.UpdateProfile(uuid, req)
	if err != nil {
		log.Errorf("update profile error: %v", err)
		c.JSON(500, err.Error())
		return
	}
	c.String(200, "success")
}

// @Summary Delete profile
// @Description Delete profile
// @Tags profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param uuid path string true "profile uuid"
// @Success 200 {string} success
// @Router /api/v1/profile/:uuid [delete]
func deleteProfile(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(400, "uuid is required")
		return
	}
	err := app.App.DBService.DeleteProfile(uuid)
	if err != nil {
		log.Errorf("delete profile error: %v", err)
		c.JSON(500, err.Error())
		return
	}
	c.String(200, "success")
}
