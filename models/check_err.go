//============================================================
// 描述: 
// 作者: Simon
// 日期: 2019/11/13 4:12 下午
//
//============================================================

package models

import (
	"github.com/pssauron/gocore/rs"
	"github.com/pssauron/log4go"
)

func CheckErr(err error ,message string) {
	if err != nil {
		log4go.Error(err)
		panic(rs.NewApiErr(999,message))
	}
}

func PanicErr(message string) {
	panic(rs.NewApiErr(999,message))
}
