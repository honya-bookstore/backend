package http

import "github.com/gin-gonic/gin"

type Router interface {
	RegisterRoutes(e *gin.Engine)
}

type RouterImpl struct{}

func ProvideRouter() *RouterImpl {
	return &RouterImpl{}
}

func (r *RouterImpl) RegisterRoutes(e *gin.Engine) {
	api := e.Group("/api")
	{
		articles := api.Group("/articles")
		{
			articles.GET("")
			articles.POST("")
			articles.GET("/:id")
			articles.PATCH("/:id")
			articles.DELETE("/:id")
		}
		books := api.Group("/books")
		{
			books.GET("")
			books.POST("")
			books.GET("/:id")
			books.PATCH("/:id")
			books.DELETE("/:id")
		}
		cart := api.Group("/cart")
		{
			cart.GET("/:id")
			cart.GET("/user/:id")
			cart.GET("/me")
			cart.POST("")
			cart.PATCH("/:id")
			cart.POST("/:id/items/item_id")
			cart.PATCH("/:id/items/item_id")
			cart.DELETE("/:id/items/item_id")

		}
		categories := api.Group("/categories")
		{
			categories.GET("")
			categories.GET("/:id")
			categories.POST("")
			categories.PATCH("/:id")
			categories.DELETE("/:id")
		}
		media := api.Group("/media")
		{
			media.GET("")
			media.GET("/:id")
			media.POST("")
			media.DELETE("/:id")
		}
		orders := api.Group("/orders")
		{
			orders.GET("")
			orders.GET("/:id")
			orders.POST("")
			orders.PUT("/:id")
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
