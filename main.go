package main

import (
	"github.com/gin-gonic/gin"
	"hutaiyi.study/gin_vue/common"
)

func main() {
	db := common.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
