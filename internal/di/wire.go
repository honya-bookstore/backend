//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"backend/config"
	"backend/internal/client"
	"backend/internal/delivery/http"

	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(
	config.NewServer,
)

var HandlerSet = wire.NewSet(
	http.ProvideArticleHandler,
	wire.Bind(
		new(http.ArticleHandler),
		new(*http.ArticleHandlerImpl),
	),
	http.ProvideBookHandler,
	wire.Bind(
		new(http.BookHandler),
		new(*http.BookHandlerImpl),
	),
	http.ProvideCartHandler,
	wire.Bind(
		new(http.CartHandler),
		new(*http.CartHandlerImpl),
	),
	http.ProvideCategoryHandler,
	wire.Bind(
		new(http.CategoryHandler),
		new(*http.CategoryHandlerImpl),
	),
	http.ProvideMediaHandler,
	wire.Bind(
		new(http.MediaHandler),
		new(*http.MediaHandlerImpl),
	),
	http.ProvideOrderHandler,
	wire.Bind(
		new(http.OrderHandler),
		new(*http.OrderHandlerImpl),
	),
	http.ProvideAuthHandler,
	wire.Bind(
		new(http.AuthHandler),
		new(*http.AuthHandlerImpl),
	),
)

var RouterSet = wire.NewSet(
	http.ProvideRouter,
	wire.Bind(
		new(http.Router),
		new(*http.RouterImpl),
	),
)

var ClientSet = wire.NewSet(
	client.NewGin,
)

func InitializeServer(ctx context.Context) *http.Server {
	wire.Build(
		ConfigSet,
		HandlerSet,
		RouterSet,
		ClientSet,
		http.NewServer,
	)
	return nil
}
