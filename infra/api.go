package infra

import (
	"merchant-service/auth"
	"merchant-service/infra/config"
	"merchant-service/product"

	"merchant-service/handler"
	"merchant-service/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.Connection()

	authService    = auth.NewService()
	userService    = user.NewService(userRepository)
	userRepository = user.NewRepository(DB)
	userHandler    = handler.NewUserHandler(userService, authService)

	productRepository = product.NewRepository(DB)
	productService    = product.NewProductService(productRepository)
	productHandler    = handler.NewProductHandler(productService, authService)
)

func RegisterApi(r *gin.Engine) {
	api := r.Group("/api")
	{
		// user
		api.POST("/user/register", userHandler.RegisterUserHandler)
		api.POST("/user/login", userHandler.LoginUserHandler)
		api.GET("/users", handler.Middleware(userService, authService), handler.AdminMiddleware(userRepository), userHandler.ShowAllUserHandler)
		api.GET("/users/:id", handler.Middleware(userService, authService), userHandler.GetUserByIDHandler)
		api.PUT("/users/:id", handler.Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
		api.DELETE("/users/:id", handler.Middleware(userService, authService), userHandler.DeleteUserByIDHandler)
		api.POST("/users/outlet", handler.Middleware(userService, authService), userHandler.CreateOutletUserHandler)
		api.GET("/users/outlet", userHandler.ShowAllOutletUserHandler)

		// product
		api.POST("/product", handler.Middleware(userService, authService), productHandler.CreateProductHandler)
		api.POST("/product/display-image", handler.Middleware(userService, authService), productHandler.CreateDisplayImageProduct)
		api.GET("/product/:outlet_id", handler.Middleware(userService, authService), productHandler.GetProductOutletByIDHandler)
		api.GET("/product", productHandler.ShowAllProductHandler)
		api.GET("/product/detail/:id", productHandler.GetProductByIDHandler)
		api.PUT("/product/:id", handler.Middleware(userService, authService), productHandler.UpdateProductByIDHandler)
		api.DELETE("/product/:id", handler.Middleware(userService, authService), productHandler.DeleteProductByIDHandler)
	}

}
