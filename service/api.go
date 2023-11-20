package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func insertDeploymentIntoDb(name string) {
	deployment := &Deployment{
		SiteName:         name,
		Theme:            "blog-1",
		ThemeOwner:       "serranoio",
		FrontendProvider: "netlify",
		BackendProvider:  "fly.io",
	}

	Db.Save(&deployment)
}

func GoatDeploy(r *gin.Engine) {
	r.GET("/goat/:name", func(c *gin.Context) {
		name := c.Param("name")

		if len(name) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid request",
				"message": "u put name of 0 length lol",
			})
		}

		insertDeploymentIntoDb(name)
		err := goated(name)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid request",
				"message": "u done f'ed up",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
}

func FetchAllDeploys(r *gin.Engine) {
	r.GET("/fetch/all", func(c *gin.Context) {

		deployments := []Deployment{}

		Db.Select("SiteName").Find(&deployments)

		c.JSON(http.StatusOK, gin.H{
			"deployments": deployments,
		})

	})

}

func GoatUpdate(r *gin.Engine) {
	r.GET("/update/all", func(c *gin.Context) {
		deployments := []Deployment{}

		Db.Select("SiteName").Find(&deployments)

		for _, deployment := range deployments {
			updateFrontend(deployment.SiteName, "blog-1")
		}

		c.JSON(http.StatusOK, gin.H{
			"deployments": deployments,
		})
	})
}
