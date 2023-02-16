package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/hermanschaaf/cq-source-xkcd/internal/xkcd"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger zerolog.Logger
	XKCD   *xkcd.Client
}

func (c *Client) ID() string {
	return "xkcd"
}

func New(_ context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	c, err := xkcd.NewClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		Logger: logger,
		XKCD:   c,
	}, nil
}
