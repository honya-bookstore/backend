package http

import "github.com/gin-gonic/gin"

type Router interface {
	RegisterRoutes(e *gin.Engine)
}

type RouterImpl struct {
	// articleHandler  ArticleHandler
	bookHandler       BookHandler
	cartHandler       CartHandler
	categoryHandler   CategoryHandler
	mediaHandler      MediaHandler
	orderHandler      OrderHandler
	authMiddleware    AuthMiddleware
	loggingMiddleware LoggingMiddleware
	metricMiddleware  MetricMiddleware
	roleMiddleware    RoleMiddleware
}

func ProvideRouter(
	// articleHandler ArticleHandler,
	bookHandler BookHandler,
	cartHandler CartHandler,
	categoryHandler CategoryHandler,
	mediaHandler MediaHandler,
	orderHandler OrderHandler,
	authMiddleware AuthMiddleware,
	loggingMiddleware LoggingMiddleware,
	metricMiddleware MetricMiddleware,
	roleMiddleware RoleMiddleware,
) *RouterImpl {
	return &RouterImpl{
		// articleHandler:  articleHandler,
		bookHandler:       bookHandler,
		cartHandler:       cartHandler,
		categoryHandler:   categoryHandler,
		mediaHandler:      mediaHandler,
		orderHandler:      orderHandler,
		authMiddleware:    authMiddleware,
		loggingMiddleware: loggingMiddleware,
		metricMiddleware:  metricMiddleware,
		roleMiddleware:    roleMiddleware,
	}
}

func (r *RouterImpl) RegisterRoutes(
	e *gin.Engine,
) {
	e.Use(r.loggingMiddleware.Handler())
	e.Use(r.metricMiddleware.Handler())
	api := e.Group("/api")
	{
		// articles := api.Group("/articles")
		// {
		// 	articles.GET("", r.articleHandler.List)
		// 	articles.POST("", r.articleHandler.Create)
		// 	articles.GET("/:id", r.articleHandler.Get)
		// 	articles.PATCH("/:id", r.articleHandler.Update)
		// 	articles.DELETE("/:id", r.articleHandler.Delete)
		// }
		books := api.Group("/books")
		{
			books.GET("", r.bookHandler.List)
			books.POST("", r.bookHandler.Create)
			books.GET("/:id", r.bookHandler.Get)
			books.PATCH("/:id", r.bookHandler.Update)
			books.DELETE("/:id", r.bookHandler.Delete)
		}
		cart := api.Group("/cart")
		{
			cart.GET("/:id", r.cartHandler.Get)
			cart.GET("/user/:user_id", r.cartHandler.GetByUser)
			cart.GET("/me", r.authMiddleware.Handler(), r.cartHandler.GetMine)
			cart.POST("", r.cartHandler.Create)
			cart.POST("/:id/items", r.cartHandler.CreateItem)
			cart.PATCH("/:id/items/:item_id", r.cartHandler.UpdateItem)
			cart.DELETE("/:id/items/:item_id", r.cartHandler.DeleteItem)
		}
		categories := api.Group("/categories")
		{
			categories.GET("", r.categoryHandler.List)
			categories.GET("/:slug", r.categoryHandler.GetBySlug)
			categories.POST("", r.categoryHandler.Create)
			categories.PATCH("/:id", r.categoryHandler.Update)
			categories.DELETE("/:id", r.categoryHandler.Delete)
		}
		media := api.Group("/media")
		{
			media.GET("", r.mediaHandler.List)
			media.GET("/:id", r.mediaHandler.Get)
			media.POST("", r.mediaHandler.Create)
			media.DELETE("/:id", r.mediaHandler.Delete)
			media.GET("/images/upload-url", r.mediaHandler.GetUploadImageURL)
			media.GET("/images/:image-id/delete-url/", r.mediaHandler.GetDeleteImageURL)
		}
		orders := api.Group("/orders")
		{
			orders.GET("", r.orderHandler.List)
			orders.GET("/:id", r.orderHandler.Get)
			orders.POST("", r.orderHandler.Create)
			orders.PUT("/:id", r.orderHandler.Update)
			orders.GET("/vnpay/ipn", r.orderHandler.VerifyVNPayIPN)
		}
		// reviews := api.Group("/reviews")
		// {
		// 	reviews.GET("")
		// 	reviews.POST("")
		// 	reviews.GET("/:id")
		// 	reviews.PATCH("/:id")
		// 	reviews.DELETE("/:id")
		// }
	}
}
