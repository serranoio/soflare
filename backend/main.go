package main

import (
	"backend/migration"

	"os"

	"github.com/charmbracelet/log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

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
		// collection, err := app.Dao().FindCollectionByNameOrId("_admins")

		// if err != nil {

		// 	log.Fatal(err)
		// }

		// log.Info(collection)

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
