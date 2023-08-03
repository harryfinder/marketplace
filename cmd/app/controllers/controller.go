package controllers

import (
	"context"
	"marketplace/internal/config"
)

type Controller interface {
	Serve(context.Context, string, config.Config) error
	Shutdown(context.Context) error
}
