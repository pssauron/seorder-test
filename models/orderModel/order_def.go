//============================================================
// 描述: 
// 作者: Simon
// 日期: 2019/11/13 3:55 下午
//
//============================================================

package orderModel

import "github.com/pssauron/gocore/libs"

type SEOrderDTO struct {
	FInterID	 libs.Int 	`db:"FInterID" json:"fInterId" primary:"true"`	//主表ID 主键
	FBillNo  		libs.String `db:"FBillNo" json:"fBillNo"`		//单据号
	FCurrencyID libs.Int 	`db:"FCurrencyID" json:"fCurrencyId"`	//币别
	FCustID 		libs.Int `db:"FCustID" json:"fCustId"`		//客户ID
	FDate				libs.Time `db:"FDate" json:"fDate"`		//制单日
	FBillerID		libs.Int 	`db:"FBillerID" json:"fBillerId"`	//制单人
	FCheckerID	libs.Int 	`db:"FCheckerID" json:"fCheckerId"`	//审核人
	FNote				libs.String `db:"FNote" json:"fNote"`	//备注
	FStatus			libs.Int	`db:"FStatus" json:"fStatus"`	//审核状态
}

func (SEOrderDTO) TableName() string {
	return "SEOrder"
}

type SEOrderEntryDTO struct {
	FInterID	  libs.Int 	`db:"FInterID" json:"fInterId"`		//单据ID
	FEntryID  	libs.Int 	`db:"FEntryID" json:"fEntryId"`		//单据行号
	FItemID			libs.Int 	`db:"FItemID" json:"fItemId"`		//物料ID
	FQty			 libs.Float `db:"FQty" json:"fQty"`			//物料数量
	FPrice		 libs.Float	`db:"FPrice" json:"fPrice"`	//单价
	FAmount    libs.Float `db:"FAmount" json:"fAmount"`	//合计金额
}

func (SEOrderEntryDTO) TableName() string {
	return "SEOrderEntry"
}



type CreateReq struct {
	*SEOrderDTO
	Entry []CreateEntryReq	`json:"entry"`
}


type CreateEntryReq struct {
	FItemID			libs.Int 	`db:"FItemID" json:"fItemId"`		//物料ID
	FQty			 libs.Float `db:"FQty" json:"fQty"`			//物料数量
	FPrice		 libs.Float	`db:"FPrice" json:"fPrice"`	//单价
}