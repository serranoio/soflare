package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "yszhz50yisvr2sb",
			"created": "2023-11-17 00:52:07.268Z",
			"updated": "2023-11-17 00:52:07.268Z",
			"name": "images",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "btjyurrg",
					"name": "img",
					"type": "file",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"maxSize": 5242880,
						"mimeTypes": null,
						"thumbs": null,
						"protected": false
					}
				}
			],
			"indexes": [],
			"listRule": "",
			"viewRule": "",
			"createRule": "@request.auth.id != ''",
			"updateRule": "@request.auth.id != ''",
			"deleteRule": "@request.auth.id != ''",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("yszhz50yisvr2sb")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}