package main

import (
	"db-operator/handler"
	"db-operator/internal"
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/prometheus/common/log"
	"strings"
)

var env string
func init() {
	flag.StringVar(&env, "env", "", "default env name")
}

func main(){
	//TODO 连接数据库
	//读配置
	paladin.Init()
	ec := internal.DBConf{
		EnvConf: make(map[string]internal.MysqlConf, 0),
	}
	if err := paladin.Get("databases.toml").UnmarshalTOML(&ec); err != nil {
		panic(err)
	}
	if err := paladin.Watch("databases.toml", &ec); err != nil {
		panic(err)
	}
	var dbMetaConf internal.MysqlConf
	checkEnvExist := false
	for key, val := range ec.EnvConf {
		if key == strings.ToUpper(env) {
			dbMetaConf = val
			checkEnvExist = true
		}
	}
	if !checkEnvExist {
		panic(errors.New("environment of database configs is not found"))
	}
	log.Info("dbconf:", dbMetaConf)
	db,err:=internal.OpenDB(ec)
	if err != nil {
		panic(err)
	}
	internal.AddDB("test",db)
	internal.SetWorkDB("test")
	router :=gin.Default()
	r1 := router.Group("/openapi/db")
	{
		r1.GET("/list/:tablename",handler.ListViewTable)
		r1.GET("/one/:tablename",handler.DetailViewTable)
		r1.POST("/new",handler.NewRow)
		r1.POST("/update",handler.UpdateRow)
		r1.POST("/del",handler.DeleteRow)
	}
	_ = router.Run(":8000")
}
