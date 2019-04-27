package routes

import (
	"distributeAccount/app/Http/Controllers/F7"
	"distributeAccount/app/Http/Controllers/Id5"
	"github.com/gin-gonic/gin"
)

func Start(router *gin.Engine) {
	//id5 router
	router.POST("/id5/getAccount", Id5.GetAccount)
	router.POST("/id5/feedbackAccount", Id5.FeedbackAccount)

	//f7 router
	router.POST("/f7/getAccount", F7.GetAccount)
}
