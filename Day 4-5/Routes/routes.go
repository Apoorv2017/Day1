package Routes

import(
	"project/Config"
	"project/Models"
	"project/Controllers"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  	"github.com/gin-gonic/gin"
	"project/Middleware"
	"fmt"
)

var err error
var router *gin.Engine

func Initialize(){
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Order{})
	
	router := gin.Default()
	grp:=router.Group("/")
	{	
		grp.GET("api-key-retailer", Middleware.GetJwtRetailer)
		grp.GET("api-key-client", Middleware.GetJwtClient)
		grp1:=grp.Group("retailer/")
		grp1.Use(Middleware.ValidateJWT())
		{
			grp1.POST("add_product", Controllers.CreatProduct)
			grp1.GET("get_product/:id", Controllers.GetProductById)
			grp1.GET("get_all_products", Controllers.GetAllProducts)
			grp1.PATCH("update_product/:id", Controllers.UpdateProduct)
			grp1.GET("get_all_orders", Controllers.GetAllOrders)
		}
		grp2:=grp.Group("customer/")
		grp2.Use(Middleware.ValidateJWT())
		{
			grp2.POST("place_order", Controllers.PlaceOrder)
			grp2.GET("get_order/:id", Controllers.GetOrderById)
		}
	}
	router.Run()

}