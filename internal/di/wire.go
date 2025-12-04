//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"backend/config"
	"backend/internal/application"
	"backend/internal/client"
	"backend/internal/delivery/http"

	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(
	config.NewServer,
)

var ApplicationSet = wire.NewSet(
	application.ProvideBook,
	wire.Bind(
		new(http.BookApplication),
		new(*application.Book),
	),
	application.ProvideCategory,
	wire.Bind(
		new(http.CategoryApplication),
		new(*application.Category),
	),
	application.ProvideCart,
	wire.Bind(
		new(http.CartApplication),
		new(*application.Cart),
	),
	application.ProvideOrder,
	wire.Bind(
		new(http.OrderApplication),
		new(*application.Order),
	),
	application.ProvideMedia,
	wire.Bind(
		new(http.MediaApplication),
		new(*application.Media),
	),
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
		ApplicationSet,
		HandlerSet,
		RouterSet,
		ClientSet,
		http.NewServer,
	)
	return nil
}
