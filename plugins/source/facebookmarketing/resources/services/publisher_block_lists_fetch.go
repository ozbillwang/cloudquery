package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchPublisherBlockLists(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	nextPage := ""
	var err error
	for {
		var items []rest.PublisherBlockList
		items, nextPage, err = cqClient.FacebookClient.ListPublisherBlockLists(ctx, nextPage)
		if err != nil {
			return err
		}

		res <- items

		if nextPage == "" {
			break
		}
	}

	return nil
}