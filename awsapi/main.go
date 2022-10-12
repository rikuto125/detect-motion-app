//https://medium.com/eureka-engineering/multipart-file-upload-in-golang-c4a8eb15a3ee
package main

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	router := gin.Default()

	// CORSの設定
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.Static("/images", "./images")
	router.POST("/getimage", handleUpload)
	//router.Run(":3000")

	//envファイルを読み込む
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	//CERT_PATHを読み込む
	certPath := os.Getenv("CERT_PATH")
	//PRIVATE_KEY_PATHを読み込む
	privateKeyPath := os.Getenv("PRIVATE_KEY_PATH")

	router.RunTLS(":8443", certPath, privateKeyPath)
}

func handleUpload(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.Status(http.StatusNotFound)
		return
	}

	// リクエストの情報を出力する
	_, err := httputil.DumpRequest(c.Request, true)
	handleError(err)
	//logを出力すると画像のbinaryが出力されるので表示に時間がかかる
	//log.Println(string(requestDump))

	// "file"というフィールド名に一致する最初のファイルが返却される
	// マルチパートフォームのデータはパースされていない場合ここでパースされる
	formFile, _, err := c.Request.FormFile("image")
	handleError(err)
	defer formFile.Close()

	// データを保存するファイルを開く
	filename := fmt.Sprintf("uploaded_%d.png", time.Now().UnixNano())
	//fmt.Println("filename:", filename)

	saveFile, err := os.Create("./images/" + filename)
	handleError(err)
	defer saveFile.Close()

	// ファイルにデータを書き込む
	_, err = io.Copy(saveFile, formFile)
	handleError(err)

	c.Status(http.StatusCreated)

	postMessageToLine()
	postImageToLine(filename)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//messageと画像を送信する
func postMessageToLine() {
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

}

func postImageToLine(dst string) {
	url := "https://api.line.me/v2/bot/message/broadcast"

	//現在のディレクトのpathを取得
	dir, _ := os.Getwd()
	fmt.Println(dir)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	//LINE_BOT_CHANNEL_ACCESS_TOKENを読み込む
	channelAccessToken := os.Getenv("LINE_BOT_CHANNEL_ACCESS_TOKEN")

	fmt.Printf("channelAccessToken: %s ", channelAccessToken)

	//1秒停止
	time.Sleep(1 * time.Second)

	//"https://techguild-test.ddns.net:8443/images/a.png"の画像をLineApiに送信

	// 送信用のデータ
	// messageの中にtype,textの配列を追加すれば一度に複数のメッセージを送信できます。(最大件数5)

	url = os.Getenv("URL")

	data := map[string][]map[string]string{
		"messages": {
			{
				"type":               "image",
				"originalContentUrl": url + dst,
				"previewImageUrl":    url + dst,
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
	fmt.Println(resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
