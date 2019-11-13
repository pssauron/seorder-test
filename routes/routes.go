//============================================================
// 描述: 
// 作者: Simon
// 日期: 2019/11/13 3:41 下午
//
//============================================================

package routes

import (
	"github.com/labstack/echo"
	"github.com/pssauron/gocore/rs"
	"github.com/seorder-test/apis/OrderApi"
	"github.com/seorder-test/middlewares"
	"net/http"
)

func InitHandlers(e *echo.Echo) {
	e.HTTPErrorHandler = ApiErrHandler
	e.Use(middlewares.Recover())

	//seorder
	initSeorder(e)


}



//	Insert  新增一张销售订单
//	UPdate	修改销售订单
//	QueryPage	分页查询订单列表
//  Del		删除一张销售订单
func initSeorder(e *echo.Echo) {

	g := e.Group("/api/order")

	g.POST("/create",OrderApi.Create)
	g.POST("/update",OrderApi.Update)
	g.POST("/del",OrderApi.Del)
	g.GET("/queryPage",OrderApi.QueryPage)

}




func ApiErrHandler(e error, c echo.Context) {

	if err, ok := e.(*rs.ApiErr); ok {
		c.JSON(http.StatusOK, err)
	} else if err, ok := e.(*echo.HTTPError); ok {
		c.JSON(err.Code, err.Message)
	} else {
		c.JSON(http.StatusInternalServerError, e.Error())
	}

}