//============================================================
// 描述: 
// 作者: Simon
// 日期: 2019/11/13 3:28 下午
//
//============================================================

package cfgs

import (
	"github.com/pssauron/gocore/stores"
	"github.com/pssauron/log4go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Cfg *Conf

type Conf struct {
	MssqlConf 	*stores.StoreConf  `yaml:"mssql"`
}

func  InitCfg(path string) {
	 bs,err := ioutil.ReadFile(path)

	 if err != nil {
	 	log4go.Error(err)
	 	panic("配置文件初始化异常")
	 }

	 err = yaml.Unmarshal(bs,&Cfg)

	 if err != nil {
	 	log4go.Error(err)
	 	panic("配置文件序列化异常")
	 }

}