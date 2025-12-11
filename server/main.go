package main

import (
	"fmt"
	"log"
	"net/http"
	_ "project_1/db"

	"github.com/gin-gonic/gin"
)

// 요청 JSON 구조체 정의

// https://kakaobusiness.gitbook.io/main/tool/chatbot/skill_guide/answer_json_format?utm_source=chatgpt.com
type KakaoRequest struct {
	Bot         BotInfo       `json:"bot"`
	Intent      IntentInfo    `json:"intent"`
	Action      ActionInfo    `json:"action"`
	UserRequest UserRequest   `json:"userRequest"`
	Contexts    []interface{} `json:"contexts"`
	Flow        FlowInfo      `json:"flow"`
}

/* --- Bot --- */
type BotInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

/* --- Intent --- */
type IntentInfo struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Extra IntentExtra `json:"extra"`
}

type IntentExtra struct {
	Reason IntentReason `json:"reason"`
}

type IntentReason struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/* --- Action --- */
type ActionInfo struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Params       map[string]interface{} `json:"params"`
	DetailParams map[string]interface{} `json:"detailParams"`
	ClientExtra  map[string]interface{} `json:"clientExtra"`
}

/* --- User Request --- */
type UserRequest struct {
	Block     BlockInfo              `json:"block"`
	User      UserInfo               `json:"user"`
	Utterance string                 `json:"utterance"`
	Params    map[string]interface{} `json:"params"`
	Lang      string                 `json:"lang"`
	Timezone  string                 `json:"timezone"`
}

type BlockInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserInfo struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Properties UserProperties `json:"properties"`
}

type UserProperties struct {
	BotUserKey          string `json:"botUserKey"`
	IsFriend            bool   `json:"isFriend"`
	PlusfriendUserKey   string `json:"plusfriendUserKey"`
	Bot_User_Key        string `json:"bot_user_key"`
	Plusfriend_User_Key string `json:"plusfriend_user_key"`
}

/* --- Flow --- */
type FlowInfo struct {
	LastBlock BlockInfo `json:"lastBlock"`
	Trigger   Trigger   `json:"trigger"`
}

type Trigger struct {
	Type string `json:"type"`
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
		var req KakaoRequest

		// body, err := io.ReadAll(c.Request.Body)
		// if err != nil {
		// 	log.Println("err: ", err.Error())
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		// 	return
		// }
		// log.Println("raw body:", string(body))

		// JSON 바인딩
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Println("err: ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 이미지도 받을 수 있음
		// https://talk.kakaocdn.net/dna/cK30ot/bl8xydVJsir/pKCAdi4a1SYvakIo2uGPcf/i_ebcab7a98468.jpg?credential=zf3biCPbmWRjbqf40YGePFLewdou7TIK&expires=1860069693&signature=kECa1AGOWLxHXchNeXXktPZvMe4%3D
		// utterance := req.UserRequest.Utterance
		// fmt.Println(utterance)                 // 콘솔 출력
		// fmt.Println(req.UserRequest.User.ID)   // 콘솔 출력
		// fmt.Println(req.UserRequest.User.Type) // 콘솔 출력

		c.JSON(http.StatusOK, gin.H{
			"version": "2.0",
			"template": gin.H{
				"outputs": []gin.H{
					{
						"simpleText": gin.H{
							"text": fmt.Sprintf("[봇] %s", req.UserRequest.Utterance),
							// "text": fmt.Sprintf("[봇] %s / %s\n%s", req.UserRequest.User.ID, req.UserRequest.User.Type, utterance),
						},
						// "simpleImage": gin.H{
						// 	"imageUrl": "https://t1.kakaocdn.net/kakaocorp/kakaocorp/admin/65626a08017800001.png",
						// },
					},
				},
				// 상담톡
				// https://kakaobusiness.gitbook.io/main/ad/cstalk
				// "conversation": gin.H{
				// 	"type":        "TRANSFER",
				// 	"clientExtra": gin.H{
				// 		// "service": gin.H{
				// 		// 	"user": gin.H{
				// 		// 		"id": req.UserRequest.User.ID,
				// 		// 	},
				// 		// },
				// 	},
				// },
			},
		})

	})

	r.Run()
}
