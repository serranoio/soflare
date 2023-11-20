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
			"id": "23mxp9lp7klkv29",
			"created": "2023-11-18 20:18:53.306Z",
			"updated": "2023-11-18 20:18:53.306Z",
			"name": "blog",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "m0wlcpls",
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
					"id": "rz9zgaz2",
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
					"id": "vuveo74s",
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
					"id": "oa2t55i4",
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
					"id": "vh17bbap",
					"name": "md",
					"type": "json",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "dm1zehj7",
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

		collection, err := dao.FindCollectionByNameOrId("23mxp9lp7klkv29")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
