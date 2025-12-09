package main

import (
	"fmt"
	"net/http"
	_ "project_1/db"

	"github.com/gin-gonic/gin"
)

// 요청 JSON 구조체 정의
type UserRequest struct {
	UserRequest struct {
		Utterance string `json:"utterance"`
	} `json:"userRequest"`
}

// 참고 url
// https://wikidocs.net/280514

// brew install ngrok/ngrok/ngrok
// ngrok config add-authtoken ${token}
// https://dashboard.ngrok.com/get-started/your-authtoken
// ngrok http 8080

// https://6e9bd08f4be5.ngrok-free.app/test
func main() {
	fmt.Println("???")
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	// POST /basic
	r.POST("/basic", func(c *gin.Context) {
		var req UserRequest

		// JSON 바인딩
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		utterance := req.UserRequest.Utterance
		fmt.Println(utterance) // 콘솔 출력

		c.JSON(http.StatusOK, gin.H{
			"version": "2.0",
			"template": gin.H{
				"outputs": []gin.H{
					{
						"simpleText": gin.H{
							"text": "안녕하세요! 니가 했던말이야" + utterance,
						},
						// "simpleImage": gin.H{
						// 	"imageUrl": "https://t1.kakaocdn.net/kakaocorp/kakaocorp/admin/65626a08017800001.png",
						// },
					},
				},
			},
		})

	})

	r.Run()
}
