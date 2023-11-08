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
		ListRule:   types.Pointer("@request.auth.id = '' || @request.auth.id != ''"),
		ViewRule:   types.Pointer("@request.auth.id = '' || @request.auth.id != ''"),
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
			},
		),
	}

	contact := &models.Collection{

		Name:       "contact",
		Type:       models.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   types.Pointer("@request.auth.id != ''"),
		CreateRule: types.Pointer(""),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "user",
				Type:     schema.FieldTypeRelation,
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect:     types.Pointer(1),
					CollectionId:  "ae40239d2bc4477",
					CascadeDelete: true,
				},
			},
			&schema.SchemaField{
				Name:     "name",
				Type:     schema.FieldTypeText,
				Required: true,
				Options: &schema.TextOptions{
					Max: types.Pointer(10),
				},
			},
			&schema.SchemaField{
				Name:     "facebook",
				Type:     schema.FieldTypeText,
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect:     types.Pointer(1),
					CollectionId:  "ae40239d2bc4477",
					CascadeDelete: true,
				},
			},
			&schema.SchemaField{
				Name:     "twitter",
				Type:     schema.FieldTypeText,
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect:     types.Pointer(1),
					CollectionId:  "ae40239d2bc4477",
					CascadeDelete: true,
				},
			},
			&schema.SchemaField{
				Name:     "instagram",
				Type:     schema.FieldTypeText,
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect:     types.Pointer(1),
					CollectionId:  "ae40239d2bc4477",
					CascadeDelete: true,
				},
			},
			&schema.SchemaField{
				Name:     "youtube",
				Type:     schema.FieldTypeText,
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect:     types.Pointer(1),
					CollectionId:  "ae40239d2bc4477",
					CascadeDelete: true,
				},
			},
		),
	}

	blog := &models.Collection{

		Name:       "blog",
		Type:       models.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   types.Pointer("@request.auth.id != ''"),
		CreateRule: types.Pointer(""),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "title",
				Type:     schema.FieldTypeText,
				Required: true,
				Options: &schema.TextOptions{
					Max: types.Pointer(10),
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
				Type:     schema.FieldTypeFile,
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect:     types.Pointer(1),
					CollectionId:  "ae40239d2bc4477",
					CascadeDelete: true,
				},
			},
		),
	}

	collections = append(collections, hero)
	collections = append(collections, contact)
	collections = append(collections, blog)

	return collections
}
