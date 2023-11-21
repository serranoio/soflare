//

package main

import (
	"os"
	"path"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func deployFrontend(name string, theme string) error {
	err := FetchFrontend(theme, name)
	if err != nil {
		return err
	}
	err = DeployToNetlify(theme, name)
	if err != nil {
		return err
	}

	return nil
}

func goated(name string) error {

	const theme = "blog-1"

	err := deployFrontend(name, theme)
	if err != nil {
		return err
	}

	err = DeployFly(name)

	return err
}

var Db *gorm.DB

type Deployment struct {
	gorm.Model
	SiteName         string `json:"sitename"`
	Theme            string `json:"theme"`
	ThemeOwner       string `json:"ThemeOwner"`
	FrontendProvider string `json:"frontendProvider"`
	BackendProvider  string `json:"backendProvider"`
}

func initDatabase() error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	pathd := path.Join(dir, "database", "database.db")
	// if file already existed, don't add metrics
	if _, err := os.Stat(pathd); os.IsNotExist(err) {
		os.Create(pathd)
	}

	Db, err = gorm.Open(sqlite.Open(pathd), &gorm.Config{})

	if err != nil {
		return err
	}

	Db.AutoMigrate(&Deployment{})

	return nil
}

func main() {
	initDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// Access-Control-Allow-Origin
		AllowMethods:     []string{"PUT", "GET", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
		AllowFiles:       true,
		MaxAge:           12 * time.Hour,
	}))

	GoatDeploy(r)
	FetchAllDeploys(r)
	// GoatUpdate(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
