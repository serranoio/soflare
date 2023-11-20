package main

import (
	"backend/migration"
	"strings"

	"os"

	"github.com/charmbracelet/log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func check(app *pocketbase.PocketBase) (bool, *models.Collection) {
	heroCollection, _ := app.Dao().FindCollectionByNameOrId("hero")

	records := []*models.Record{}

	app.Dao().RecordQuery(heroCollection).All(&records)

	if len(records) > 0 {
		return false, heroCollection
	}

	return true, heroCollection
}

func createAdmin(app *pocketbase.PocketBase) {
	admin := &models.Admin{}
	admin.Email = "admin@gmail.com"
	admin.SetPassword("1234567890")

	err := app.Dao().SaveAdmin(admin)

	if err != nil {
		log.Print("Could not set admin")
		err = nil
	}
}
func initial(app *pocketbase.PocketBase) {
	initial, heroCollection := check(app)

	if !initial {
		return
	}

	// createAdmin(app)

	heroRecord := models.NewRecord(heroCollection)
	heroRecord.Set("title", "Title")
	heroRecord.Set("subtitle", "subtitle")
	heroRecord.Set("description", "description")
	heroRecord.Set("second-description", "second-description")
	heroRecord.Set("pic", " ")

	if err := app.Dao().SaveRecord(heroRecord); err != nil {
		log.Debug(err)
	}

	contactCollection, _ := app.Dao().FindCollectionByNameOrId("contact")
	contactRecord := models.NewRecord(contactCollection)

	contactRecord.Set("name", "Your Name")
	contactRecord.Set("email", "Your Email")
	contactRecord.Set("facebook", " ")
	contactRecord.Set("twitter", " ")
	contactRecord.Set("instagram", " ")
	contactRecord.Set("youtube", " ")
	contactRecord.Set("github", " ")

	if err := app.Dao().SaveRecord(contactRecord); err != nil {
		log.Debug(err)
	}

	metadataCollection, _ := app.Dao().FindCollectionByNameOrId("admin-controls")
	metadataRecord := models.NewRecord(metadataCollection)

	metadataRecord.Set("accents", "red")
	metadataRecord.Set("design", "1")

	if err := app.Dao().SaveRecord(metadataRecord); err != nil {
		log.Debug(err)
	}

}

func migrate(app *pocketbase.PocketBase) {
	collections := migration.CreateCollections()

	for _, collection := range collections {
		if err := app.Dao().SaveCollection(collection); err != nil {
			// log.Fatal(err)
		}
	}

}

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		migrate(app)
		initial(app)
		// collection, err := app.Dao().FindCollectionByNameOrId("_admins")

		// if err != nil {

		// 	log.Fatal(err)
		// }

		// log.Info(collection)

		return nil
	})

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
	// createAdmin()
}
