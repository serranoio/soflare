package migration

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func CreateCollections() []*models.Collection {
	collections := []*models.Collection{}

	hero := &models.Collection{

		Name:       "hero",
		Type:       models.CollectionTypeBase,
		ListRule:   types.Pointer(""),
		ViewRule:   types.Pointer(""),
		CreateRule: types.Pointer(""),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "title",
				Type:     schema.FieldTypeText,
				Required: true,
			},
			&schema.SchemaField{
				Name:     "subtitle",
				Type:     schema.FieldTypeText,
				Required: false,
			},
			&schema.SchemaField{
				Name:     "description",
				Type:     schema.FieldTypeText,
				Required: true,
			},
			&schema.SchemaField{
				Name:     "second-description",
				Type:     schema.FieldTypeText,
				Required: false,
			},
			&schema.SchemaField{
				Name:     "pic",
				Type:     schema.FieldTypeFile,
				Required: false,
				Options: &schema.FileOptions{
					MaxSelect: 1,
					MaxSize:   5242880,
				},
			},
		),
	}

	contact := &models.Collection{

		Name:       "contact",
		Type:       models.CollectionTypeBase,
		ListRule:   types.Pointer(""),
		ViewRule:   types.Pointer(""),
		CreateRule: types.Pointer(""),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "name",
				Type:     schema.FieldTypeText,
				Required: true,
			},
			&schema.SchemaField{
				Name:     "email",
				Type:     schema.FieldTypeText,
				Required: false,
			},
			&schema.SchemaField{
				Name:     "facebook",
				Type:     schema.FieldTypeText,
				Required: false,
			},
			&schema.SchemaField{
				Name:     "twitter",
				Type:     schema.FieldTypeText,
				Required: false,
			},
			&schema.SchemaField{
				Name:     "instagram",
				Type:     schema.FieldTypeText,
				Required: false,
			},
			&schema.SchemaField{
				Name:     "youtube",
				Type:     schema.FieldTypeText,
				Required: false,
			},
			&schema.SchemaField{
				Name:     "github",
				Type:     schema.FieldTypeText,
				Required: false,
			},
		),
	}

	blog := &models.Collection{

		Name:       "blog",
		Type:       models.CollectionTypeBase,
		ListRule:   types.Pointer(""),
		ViewRule:   types.Pointer(""),
		CreateRule: types.Pointer("@request.auth.id != ''"),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: types.Pointer(""),
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "title",
				Type:     schema.FieldTypeText,
				Required: true,
				Options: &schema.TextOptions{
					Max: types.Pointer(50),
				},
			},
			&schema.SchemaField{
				Name:     "subtitle",
				Type:     schema.FieldTypeText,
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect:     types.Pointer(1),
					CollectionId:  "ae40239d2bc4477",
					CascadeDelete: true,
				},
			},
			&schema.SchemaField{
				Name:     "author",
				Type:     schema.FieldTypeText,
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect:     types.Pointer(1),
					CollectionId:  "ae40239d2bc4477",
					CascadeDelete: true,
				},
			},
			&schema.SchemaField{
				Name:     "date",
				Type:     schema.FieldTypeDate,
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect:     types.Pointer(1),
					CollectionId:  "ae40239d2bc4477",
					CascadeDelete: true,
				},
			},

			&schema.SchemaField{
				Name:     "md",
				Type:     "json",
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect:     types.Pointer(1),
					CollectionId:  "ae40239d2bc4477",
					CascadeDelete: true,
				},
			},
			&schema.SchemaField{
				Name:     "titleImg",
				Type:     schema.FieldTypeFile,
				Required: true,
				Options: &schema.FileOptions{
					MaxSelect: 1,
					MaxSize:   5242880,
				},
			},
		),
	}

	metadata := &models.Collection{

		Name:       "admin-controls",
		Type:       models.CollectionTypeBase,
		ListRule:   types.Pointer(""),
		ViewRule:   types.Pointer(""),
		CreateRule: types.Pointer("@request.auth.id != ''"),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: types.Pointer(""),
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "accents",
				Type:     schema.FieldTypeText,
				Required: true,
			},
			&schema.SchemaField{
				Name:     "design",
				Type:     schema.FieldTypeText,
				Required: true,
			},
		),
	}

	images := &models.Collection{

		Name:       "images",
		Type:       models.CollectionTypeBase,
		ListRule:   types.Pointer(""),
		ViewRule:   types.Pointer(""),
		CreateRule: types.Pointer("@request.auth.id != ''"),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: types.Pointer("@request.auth.id != ''"),
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "img",
				Type:     schema.FieldTypeFile,
				Required: true,
				Options: &schema.FileOptions{
					MaxSelect: 1,
					MaxSize:   5242880,
				},
			},
		),
		// Schema: schema.NewSchema(
		// 	&schema.SchemaField{
		// 		Name:     "img",
		// 		Type:     schema.FieldNameId,
		// 		Required: true,
		// 	},
		// ),
	}

	collections = append(collections, hero)
	collections = append(collections, contact)
	collections = append(collections, blog)
	collections = append(collections, metadata)
	collections = append(collections, images)

	return collections
}
