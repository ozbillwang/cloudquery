package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Adlabels() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_adlabels",
		Resolver:    fetchAdlabels,
		Transform:   transformers.TransformWithStruct(&rest.Adlabel{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-label#Reading",
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Resolver: client.ResolveAccountId,
				Type:     schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}