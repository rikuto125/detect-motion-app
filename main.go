package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/rikut0125/detect-motion-app/def"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

//raspi環境
func main() {
	for {
		//1秒停止
		time.Sleep(1 * time.Second)

		path := def.PointCameraDetectMotion()

		//pathの文字列の/より後ろの文字列を取得
		filename := path[11:]

		//画像を送信
		//envからurlを取得
		url := os.Getenv("URL")

		fieldname := "image"
		//現在のディレクトのpathを取得
		dir, _ := os.Getwd()
		fmt.Println(dir)

		file, err := os.Open(path)
		handleError(err)

		// リクエストボディのデータを受け取るio.Writerを生成する。
		body := &bytes.Buffer{}

		// データのmultipartエンコーディングを管理するmultipart.Writerを生成する。
		// ランダムなbase-16バウンダリが生成される。
		mw := multipart.NewWriter(body)

		// ファイルに使うパートを生成する。
		// ヘッダ以外はデータは書き込まれない。
		// fieldnameとfilenameの値がヘッダに含められる。
		// ファイルデータを書き込むio.Writerが返却される。
		fw, err := mw.CreateFormFile(fieldname, filename)

		// fwで作ったパートにファイルのデータを書き込む
		_, err = io.Copy(fw, file)
		handleError(err)

		// リクエストのContent-Typeヘッダに使う値を取得する（バウンダリを含む）
		contentType := mw.FormDataContentType()

		// 書き込みが終わったので最終のバウンダリを入れる
		err = mw.Close()
		handleError(err)

		// contentTypeとbodyを使ってリクエストを送信する
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Post(url, contentType, body)
		//resp, err := http.Post(url, contentType, body)
		handleError(err)
		println(resp)

		err = resp.Body.Close()
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
