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
			"id": "70qqovzikdrvgfx",
			"created": "2023-11-16 16:24:53.376Z",
			"updated": "2023-11-16 16:24:53.376Z",
			"name": "blog",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "pxacsbaa",
					"name": "title",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": 50,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "cn7cunt8",
					"name": "subtitle",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "mnm0syix",
					"name": "author",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "askgfcmu",
					"name": "date",
					"type": "date",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "a6nl6ob6",
					"name": "md",
					"type": "json",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "79jqc9sw",
					"name": "titleImg",
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
			"deleteRule": "",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("70qqovzikdrvgfx")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
