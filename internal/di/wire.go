//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"backend/config"
	"backend/internal/application"
	"backend/internal/client"
	"backend/internal/delivery/http"
	"backend/internal/domain"
	"backend/internal/infrastructure/objectstorages3"
	"backend/internal/infrastructure/paymentservice"
	"backend/internal/infrastructure/repositorypostgres"
	"backend/internal/service"

	"backend/pkg/logger"
	"github.com/google/wire"
)

var ConfigSet = wire.NewSet(
	config.NewServer,
	logger.NewConfig,
)

var LoggerSet = wire.NewSet(
	logger.New,
)

var DbSet = wire.NewSet(
	client.NewDBConnection,
	client.NewDBQueries,
	client.NewDBTransactor,
)

var ObjectStorageSet = wire.NewSet(
	objectstorages3.ProvideMedia,
	wire.Bind(
		new(application.MediaObjectStorage),
		new(*objectstorages3.Media),
	),
)

var ServiceSet = wire.NewSet(
	service.ProvideBook,
	wire.Bind(
		new(domain.BookService),
		new(*service.Book),
	),
	service.ProvideCategory,
	wire.Bind(
		new(domain.CategoryService),
		new(*service.Category),
	),
	service.ProvideCart,
	wire.Bind(
		new(domain.CartService),
		new(*service.Cart),
	),
	service.ProvideOrder,
	wire.Bind(
		new(domain.OrderService),
		new(*service.Order),
	),
	service.ProvideMedia,
	wire.Bind(
		new(domain.MediaService),
		new(*service.Media),
	),
)

var RepositorySet = wire.NewSet(
	repositorypostgres.ProvideCategory,
	wire.Bind(
		new(domain.CategoryRepository),
		new(*repositorypostgres.Category),
	),
	repositorypostgres.ProvideBook,
	wire.Bind(
		new(domain.BookRepository),
		new(*repositorypostgres.Book),
	),
	repositorypostgres.ProvideCart,
	wire.Bind(
		new(domain.CartRepository),
		new(*repositorypostgres.Cart),
	),
	repositorypostgres.ProvideOrder,
	wire.Bind(
		new(domain.OrderRepository),
		new(*repositorypostgres.Order),
	),
	repositorypostgres.ProvideMedia,
	wire.Bind(
		new(domain.MediaRepository),
		new(*repositorypostgres.Media),
	),
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
	// http.ProvideArticleHandler,
	// wire.Bind(
	// 	new(http.ArticleHandler),
	// 	new(*http.ArticleHandlerImpl),
	// ),
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

var MiddlewareSet = wire.NewSet(
	http.ProvideAuthMiddleware,
	wire.Bind(
		new(http.AuthMiddleware),
		new(*http.AuthMiddlewareImpl),
	),
	http.ProvideLoggingMiddleware,
	wire.Bind(
		new(http.LoggingMiddleware),
		new(*http.LoggingMiddlewareImpl),
	),
	http.ProvideMetricMiddleware,
	wire.Bind(
		new(http.MetricMiddleware),
		new(*http.MetricMiddlewareImpl),
	),
	http.ProvideRoleMiddleware,
	wire.Bind(
		new(http.RoleMiddleware),
		new(*http.RoleMiddlewareImpl),
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
	client.NewValidate,
	client.NewS3,
	client.NewS3Presign,
	client.NewKeycloak,
)

var PaymentServiceSet = wire.NewSet(
	paymentservice.ProvideVNPay,
	wire.Bind(
		new(application.OrderPaymentService),
		new(*paymentservice.VNPay),
	),
)

func InitializeServer(ctx context.Context) *http.Server {
	wire.Build(
		ApplicationSet,
		ClientSet,
		ConfigSet,
		DbSet,
		LoggerSet,
		MiddlewareSet,
		HandlerSet,
		PaymentServiceSet,
		RepositorySet,
		RouterSet,
		ServiceSet,
		ObjectStorageSet,
		http.NewServer,
	)
	return nil
}
