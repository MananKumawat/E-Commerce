package Routes
import (
	"awesomeProject/Day45/Exercise/Controllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/ecommerce-api")
	{
		grp1.GET("product", Controllers.GetProducts)
		grp1.POST("product", Controllers.CreateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.PUT("product/:id", Controllers.UpdateProduct)
		grp1.DELETE("product/:id", Controllers.DeleteProduct)

		grp1.GET("order", Controllers.GetOrders)
		grp1.POST("order", Controllers.CreateOrder)
		grp1.GET("order/:id", Controllers.GetOrderByID)
		grp1.GET("orders/:customer_id", Controllers.GetOrderByCustomerID)
		grp1.PUT("order/:id", Controllers.UpdateOrder)
		grp1.DELETE("order/:id", Controllers.DeleteOrder)
		grp1.GET("transactions", Controllers.GetOrderPlaced)

	}
	return r
}