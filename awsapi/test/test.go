package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"github.com/gin-contrib/cors"
)

//LineAPIにgetImageで受け取った画像をPostする
func postImageToLine(dst string) {
	url := "https://api.line.me/v2/bot/message/broadcast"

	//imageUrl := "localhost:8085/imag/"

	// ./.envファイルを読み込む
	dir, _ := os.Getwd()
	fmt.Println(dir)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	//LINE_BOT_CHANNEL_ACCESS_TOKENを読み込む
	channelAccessToken := os.Getenv("LINE_BOT_CHANNEL_ACCESS_TOKEN")
	//1秒停止
	time.Sleep(1 * time.Second)

	//nowTime

	text := "異常あり"

	// 送信用のデータ
	// messageの中にtype,textの配列を追加すれば一度に複数のメッセージを送信できます。(最大件数5)
	data := map[string][]map[string]string{
		"messages": {
			{
				"type": "text",
				"text": text,
			},
		},
	}

	// JSONに変換
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// リクエストを作成
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
		return
	}

	// ヘッダーを設定
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+channelAccessToken)

	// リクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// レスポンスを表示
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

	//画像を送信

	// 送信用のデータ
	// messageの中にtype,textの配列を追加すれば一度に複数のメッセージを送信できます。(最大件数5)
	data = map[string][]map[string]string{
		"messages": {
			{
				"type":               "image",
				"originalContentUrl": "https://techguild-test.ddns.net:10443/" + dst,
			},
		},
	}

	// JSONに変換
	jsonStr, err = json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// リクエストを作成
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
		return
	}

	// ヘッダーを設定
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+channelAccessToken)

	// リクエストを送信
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// レスポンスを表示
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

// raspiから送られてきた画像を受け取り、imagesディレクトリに保存する
func getImage(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	fmt.Println(files)
	for _, file := range files {
		fmt.Println(file.Filename)
	}

	// Save files to specific dst.
	for _, file := range files {
		// Upload the file to specific dst.
		dst := "images/" + file.Filename
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}
	}

	dst := "images/" + files[0].Filename
	postImageToLine(dst)
}

// gin web framework
func main() {
	r := gin.Default()
	
	config := cors.DefaultConfig()
    	config.AllowAllOrigins = true
    	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/getimage", getImage)

	r.Static("/images", "./images")

	r.RunTLS(":8443", "/etc/letsencrypt/live/techguild-test.ddns.net/cert.pem", "/etc/letsencrypt/live/techguild-test.ddns.net/privkey.pem")
}
