package main

import (
	"context"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	traq "github.com/traPtitech/go-traq"
)


func main() {
	token := os.Getenv("TRAQ_TOKEN")

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, token)

	e := echo.New()


	e.GET("/api/ubugoe/:userId", func(c echo.Context) error { 
		// フロントからは同じドメイン(https:ubugoechecker)の/apiにアクセスを飛ばす　(vite.config.tsで設定している)
		// /api →　ならバックエンドという風にする。ので変更
		userID := c.Param("userId")
		//fmt.Println(channelID)

		channelList, _, err := client.ChannelApi.GetChannels(auth).Path("gps/times/" + userID).
			Execute()
		if err != nil {
			return c.String(404, "Channel not found")
		}
		
		if len(channelList.Public) == 0{
			return c.String(404, "No channel")
		}

		channel := channelList.Public[0]
		channelID := channel.Id

		messages, _, _ := client.ChannelApi.GetMessages(auth, channelID).
			Since(time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)).
			Limit(10).
			Order("asc").
			Execute()

		messageContents := make([]string, 0, len(messages))
		for _, message := range messages {
			messageContents = append(messageContents, message.Content)
		}

		return c.JSON(200, messageContents)
	})

	e.Start(":8080")

}
