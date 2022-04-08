package firestore

import (
	"context"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

type Controller struct {
	c *datastore.Client
}

func NewController(ctx context.Context, projectID string, opts ...option.ClientOption) (*Controller, error) {
	c, err := datastore.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	return &Controller{c}, nil
}

func (ctrler *Controller) Close() error {
	return ctrler.c.Close()
}

