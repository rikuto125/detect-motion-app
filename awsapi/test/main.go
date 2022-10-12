//package main
//
//import (
//	"bytes"
//	"fmt"
//	"github.com/goccy/go-json"
//	"github.com/joho/godotenv"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"time"
//)
//
//// ラインapiに送信する関数
//
//func main() {
//	//PostXserver()
//	url := "https://api.line.me/v2/bot/message/broadcast"
//
//	//imageUrl := "localhost:8085/imag/"
//
//	// ./.envファイルを読み込む
//	dir, _ := os.Getwd()
//	fmt.Println(dir)
//
//	err := godotenv.Load("./test/.env")
//	if err != nil {
//		fmt.Println("Error loading .env file")
//	}
//
//	//LINE_BOT_CHANNEL_ACCESS_TOKENを読み込む
//	channelAccessToken := os.Getenv("LINE_BOT_CHANNEL_ACCESS_TOKEN")
//	//1秒停止
//	time.Sleep(1 * time.Second)
//
//	text := "test"
//
//	// 送信用のデータ
//	// messageの中にtype,textの配列を追加すれば一度に複数のメッセージを送信できます。(最大件数5)
//	data := map[string][]map[string]string{
//		"messages": {
//			{
//				"type": "text",
//				"text": text,
//			},
//		},
//	}
//
//	// JSONに変換
//	jsonStr, err := json.Marshal(data)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// リクエストを作成
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// ヘッダーを設定
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Authorization", "Bearer "+channelAccessToken)
//
//	// リクエストを送信
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// レスポンスを表示
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(body))
//
//	//画像を送信
//
//	// 送信用のデータ
//	// messageの中にtype,textの配列を追加すれば一度に複数のメッセージを送信できます。(最大件数5)
//	data = map[string][]map[string]string{
//		"messages": {
//			{
//				"type":               "image",
//				"originalContentUrl": "http://localhost:8085/static/b.png",
//			},
//		},
//	}
//
//	// JSONに変換
//	jsonStr, err = json.Marshal(data)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// リクエストを作成
//	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// ヘッダーを設定
//	req.Header.Set("Content-Type", "application/json")
//	req.Header.Set("Authorization", "Bearer "+channelAccessToken)
//
//	// リクエストを送信
//	client = &http.Client{}
//	resp, err = client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// レスポンスを表示
//	defer resp.Body.Close()
//
//	body, err = ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(body))
//}
