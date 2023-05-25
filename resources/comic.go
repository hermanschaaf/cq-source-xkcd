package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/hermanschaaf/cq-source-xkcd/client"
	"github.com/hermanschaaf/cq-source-xkcd/internal/xkcd"
	"golang.org/x/sync/errgroup"
)

func ComicsTable() *schema.Table {
	return &schema.Table{
		Name:      "xkcd_comics",
		Resolver:  fetchComics,
		Transform: transformers.TransformWithStruct(&xkcd.Comic{}),
	}
}

func fetchComics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	comic, err := c.XKCD.GetLatestComic(0)
	if err != nil {
		return err
	}
	res <- comic
	g := errgroup.Group{}
	g.SetLimit(10)
	for i := 1; i < comic.Num; i++ {
		i := i
		g.Go(func() error {
			comic, err := c.XKCD.GetComic(i)
			if err != nil {
				c.Logger.Error().Err(err).Msgf("failed to fetch comic %d", i)
				return err
			}
			res <- comic
			return nil
		})
	}
	return g.Wait()
}
