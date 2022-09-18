package router

import (
	"strconv"
	"time"
	"whoAmI/service/difficult"
	"whoAmI/service/easy"
	"whoAmI/service/middle"
	"whoAmI/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
创建路由
@param port 服务端口号（eg: ":8080"）
*/
func InitRouter(port string) {
	r := gin.Default()
	r.Use(cors.Default()) // 跨域

	r.GET("/hello", func(c *gin.Context) {
		wrapC := util.Gin{C: c}
		wrapC.Response(util.SUCCESS, map[string]string{
			"time": time.Now().Format("2006-01-02T15:04:05-0700"),
			"text": "hello, world",
		})
	})
	easyRouter := r.Group("/easy")
	{
		easyRouter.GET("/getList", easyGetListHandler)
	}
	middleRouter := r.Group("/middle")
	{
		middleRouter.GET("/getQuestion", middleGetQuesWithAHandler)
	}
	difficultRouter := r.Group("/difficult")
	{
		difficultRouter.GET("/getQuestion", difficultGetQuesHandler)
		difficultRouter.POST("/verify", difficultVerify)
	}

	r.Static("/static", "./static")

	r.Run(port)
}

func easyGetListHandler(c *gin.Context) {
	wrapC := util.Gin{C: c}
	picdatas := easy.GetPicDatas()
	wrapC.Response(util.SUCCESS, map[string]any{
		"picdatas": picdatas,
	})
}

func middleGetQuesWithAHandler(c *gin.Context) {
	wrapC := util.Gin{C: c}
	quesWithA := middle.GetQuestionWithA()
	wrapC.Response(util.SUCCESS, quesWithA)
}

func difficultGetQuesHandler(c *gin.Context) {
	wrapC := util.Gin{C: c}
	ques := difficult.GetQuestion()
	wrapC.Response(util.SUCCESS, ques)
}

type formDifficultVerify struct {
	Id         string `form:"id" json:"id"`
	FileName   string `json:"filename" form:"filename" binding:"required"`
	ChosenId   string `form:"chosenId" json:"chosenId" binding:"required"`
	ChosenText string `json:"chosenText" form:"chosenText" binding:"required"`
	ChosenName string `json:"chosenName" form:"chosenName" binding:"required"`
}

func difficultVerify(c *gin.Context) {
	wrapC := util.Gin{C: c}

	verifyForm := formDifficultVerify{}

	if errA := c.ShouldBind(&verifyForm); errA == nil {
		n, err := strconv.Atoi(verifyForm.ChosenId)
		if err != err {
			wrapC.Response(util.REQ_PARAM_INVALID_ERR, nil)
		}
		wrapC.Response(util.SUCCESS, map[string]bool{
			"status": difficult.Verify(
				verifyForm.Id,
				verifyForm.FileName,
				n,
				verifyForm.ChosenText,
				verifyForm.ChosenName,
			),
		})
	} else {
		wrapC.Response(util.REQ_PARAM_INVALID_ERR, nil)
	}

}
