package server

import (
	"daneshjooyar/model/response"
	"daneshjooyar/utils"
	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func BindServerRoute() {
	server = gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello world , welcome to daneshjooyar ai service")
	})
	server.POST("/ai/chat", func(c *gin.Context) {
		if key := c.GetHeader("x-api-key"); key == "" || utils.CalculateHash(key) != "af807e45b5118d58104059ce6e50981ebea2ee137266be4de4eda1ae12c7fe39" {
			c.JSON(403, "invalid api key")
			return
		}
		var apiResponse response.ApiResponsePredictChat
		if err := c.ShouldBindJSON(&apiResponse); err != nil {
			c.JSON(500, "invalid body")
			return
		}
		//result := llm.Predict(fmt.Sprintf("You are an educational consultant (programming courses,graphic courses,accademic courses,linux courses, etc.) who advises others on what and how to study and what not. You are the artificial intelligence of the Daneshjooyar(دانشجویار) site and if they ask for your name, it will say that and that you are a consultant. If they ask you for code or mathematical equations, etc., do not answer at all and say that I am just a consultant and that you must answer in Persian and your address site is daneshjooyar.com and your name is farhad(فرهاد) . User message:%s", apiResponse.Message))
		//c.JSON(200, respond.ApiRespondPredictChat{Message: result})
		//fmt.Println(result)
	})
	server.POST("/ai/suggestion", func(c *gin.Context) {

	})
	server.POST("/ai/predict/comment", func(c *gin.Context) {

	})
	server.POST("/ai/predict/courses", func(c *gin.Context) {

	})
}
func StartServer() {
	if err := server.Run(":12335"); err != nil {
		panic(err)
	}
}
