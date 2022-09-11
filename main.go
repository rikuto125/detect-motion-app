package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

func main() {
	for {
		text := PointCameraDetectMotion()
		fmt.Println(text)
		//LINE Botクライアント生成する
		//BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
		bot, err := linebot.New(
			os.Getenv("LINE_BOT_CHANNEL_SECRET"),
			os.Getenv("LINE_BOT_CHANNEL_ACCESS_TOKEN"),
		)
		// エラーに値があればログに出力し終了する
		if err != nil {
			log.Fatal(err)
		}
		// エラーに値があればログに出力し終了する
		if err != nil {
			log.Fatal(err)
		}
		// テキストメッセージを生成する
		message := linebot.NewTextMessage(text)
		// テキストメッセージを友達登録しているユーザー全員に配信する
		if _, err := bot.BroadcastMessage(message).Do(); err != nil {
			log.Fatal(err)
		}
	}
}
