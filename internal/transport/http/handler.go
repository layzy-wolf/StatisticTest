package transport

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HTTPHandler(store *gorm.DB) *gin.Engine {
	router := gin.Default()

	tr := MakeAppHandler(store)
	r := router.Group("/order")
	{
		r.GET("/book", gin.WrapH(tr.GetOrderBook))
		r.POST("/book/save", gin.WrapH(tr.SaveOrderBook))
		r.GET("/history", gin.WrapH(tr.GetOrderHistory))
		r.POST("/save", gin.WrapH(tr.SaveOrder))
	}

	return router
}
