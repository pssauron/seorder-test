//============================================================
// 描述: 
// 作者: Simon
// 日期: 2019/11/13 4:04 下午
//
//============================================================

package orderModel

import (
	"github.com/pssauron/gocore/libs"
	"github.com/seorder-test/models"
	"github.com/seorder-test/storages"
)

func CreateSEOrder(req *CreateReq) {
	q := `insert into SEOrder(FBillNo,FCurrencyID,FCustID,FDate,FBillerID,FNote,FStatus) values(?,?,?,?,?,?,?)`

	tx,err := storages.UseMssql.Begin()
	models.CheckErr(err,"数据库事务异常")


	_,err = tx.Exec(q,req.FBillNo,req.FCurrencyID,req.FCustID,req.FDate,req.FBillerID,req.FNote,0)

	models.CheckErr(err,"保存主表异常")

	var id int

	err = storages.UseMssql.Get(&id,"select ident_current('SEOrder')")
	models.CheckErr(err,"保存主表异常")
	if err != nil {
		tx.Rollback()
	}

	models.CheckErr(err,"获取插入主键异常")

	//循环保存明细表记录
	q = `insert into SEOrderEntry(FInterID,FEntryID,FItemID,FPrice,FQty,FAmount) values (?,?,?,?,?,?)`
	for index,item := range req.Entry {
		//
		entry := new(SEOrderEntryDTO)
		entry.FInterID = libs.NewInt(int(id))
		entry.FEntryID = libs.NewInt(index+1)
		entry.FItemID = item.FItemID
		entry.FPrice = item.FPrice
		entry.FQty = item.FQty
		entry.FAmount = libs.NewFloat(item.FPrice.Get() * item.FQty.Get())

		_,err := tx.Exec(q,entry.FInterID.Get(),entry.FEntryID.Get(),entry.FItemID.Get(),entry.FPrice.Get(),entry.FQty.Get(),entry.FAmount.Get())
		if err != nil {
			tx.Rollback()
		}
		models.CheckErr(err,"保存行记录异常")

	}

	tx.Commit()
}
