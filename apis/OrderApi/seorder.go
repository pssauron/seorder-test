//============================================================
// 描述: 
// 作者: Simon
// 日期: 2019/11/13 3:47 下午
//
//============================================================

package OrderApi

import (
	"github.com/labstack/echo"
	"github.com/pssauron/gocore/rs"
	"github.com/seorder-test/models"
	"github.com/seorder-test/models/orderModel"
	"net/http"
)




// 新增订单
func Create(ctx echo.Context) error {

	//1.
	req := new(orderModel.CreateReq)

	//echo json 参数病毒到req
	err := ctx.Bind(req)
	models.CheckErr(err,"参数绑定异常")


	// todo: 验证 传入的参数的完整性

	if req.FBillerID.IsEmpty()  {
		models.PanicErr("制单人不能为空")
	}


	orderModel.CreateSEOrder(req)

	return ctx.JSON(http.StatusOK,rs.NewNoDataResult())
}

func Update(ctx echo.Context) error {

	return nil
}


func QueryPage(ctx echo.Context) error {

	return nil
}

func Del(ctx echo.Context) error {

	return nil
}