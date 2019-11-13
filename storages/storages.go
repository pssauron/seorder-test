//============================================================
// 描述: 
// 作者: Simon
// 日期: 2019/11/13 3:36 下午
//
//============================================================

package storages

import "github.com/pssauron/gocore/stores"

var (
	UseMssql *stores.MSStore
)

//初始化数据操作对象
func InitStorage(conf *stores.StoreConf) {
	UseMssql = stores.NewMSStore(conf)
}
