package main

import (
	"blog-api/model"
	"github.com/freezeChen/studio-library/middle"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	engine := gin.Default()
	engine.Use(middle.CORSMiddleware())

	engine.GET("/getHomeList", getHomeList)
	err := engine.Run(":8081")
	if err != nil {
		panic(err)
	}

}

func getHomeList(ctx *gin.Context) {
	var list = make([]*model.Article, 0)

	list = append(list, &model.Article{
		Id:       1,
		Title:    "文章1",
		Detail:   "golang 内容",
		Language: "golang",
		Author:   "freeze",
		Time:     time.Now().String(),
	})

	list = append(list, &model.Article{
		Id:       2,
		Title:    "文章2",
		Detail:   "java 内容",
		Language: "Java",
		Author:   "freeze",
		Time:     time.Now().String(),
	})

	ctx.JSON(200, list)

}
