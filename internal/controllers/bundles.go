package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//	@BasePath	/api/v1
//
// TODO: update when implement
// CreateBundleController godoc
//
// CreateBundleController defines the route and Swagger annotations for create single bundle.
// @Summary      Create bundle with wallet addresses information
// @Description  Create bundle with wallet addresses information
// @Tags         bundles
// @Accept       json
// @Produce      json
// @Param 		 addresses body nil true "Bundles request"
// @Success      200 {string} string
// @Failure      400 {object} errors.APIError
// @Failure      500 {object} errors.APIError
// @Router       /api/v1/bundles [post]
func CreateBundleController(c *gin.Context, db *gorm.DB) {}

//	@BasePath	/api/v1

// GetBundlesByIdController godoc
// TODO: update when implement
// GetBundlesByIdController defines the route and Swagger annotations for fetching single bundle.
// @Summary      Fetch bundle addresses information
// @Description  Fetch bundle addresses information
// @Tags         bundles
// @Accept       json
// @Produce      json
// @Param        id path int true "Bundle ID" Format(int)
// @Success      200 {string} string
// @Failure      400 {object} errors.APIError
// @Failure      500 {object} errors.APIError
// @Router       /api/v1/bundles/:id [get]
func GetBundlesByIdController(c *gin.Context, db *gorm.DB) {}
