package main

import (
	"colloc_rest/internal/app/pkg/db"
	"colloc_rest/internal/app/pkg/order"
	pgorder "colloc_rest/internal/app/pkg/order/postgresql"
	"colloc_rest/internal/app/pkg/order_processing"
	"colloc_rest/internal/app/pkg/product"
	"colloc_rest/internal/app/pkg/product/postgresql"
	"colloc_rest/internal/app/pkg/product_order"
	pgproductorder "colloc_rest/internal/app/pkg/product_order/postgresql"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func HelloWorld(c *gin.Context) {
	fmt.Println("Hello, World!")
	c.String(http.StatusOK, "hi")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//	@title			Swagger API
//	@version		1.0
//	@description	UI for microservers.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.NewDB(ctx)
	if err != nil {
		fmt.Println("db connection error")
		return
	}
	defer database.GetPool(ctx).Close()

	productRepo := postgresql.NewProduct(database)
	orderRepo := pgorder.NewOrder(database)
	productOrderRepo := pgproductorder.NewProductOrder(database)

	orderProcessingService := order_processing.NewService(product.NewService(productRepo), product_order.NewService(productOrderRepo), order.NewService(orderRepo))

	router := gin.New()

	router.Use(CORSMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/products", orderProcessingService.GetProducts)
	router.GET("/cart", orderProcessingService.GetCartByUserId)
	//router.GET("/orders", orderProcessingService.GetOrderById)

	router.Run(":8080")
}
