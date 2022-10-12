// 参考Url
// https://algorithm.joho.info/programming/python/opencv-frame-difference-surveillance-camera-py/
//  window2．の部分のコメントアウトを外すと差分の動きがわかる.(window1をコメントアウトすること)

package def

import (
	"fmt"
	"gocv.io/x/gocv"
	"time"
)

func PointCameraDetectMotion() string {
	//閾値 これより大きい差分があったら動いたと判定する　この値は調整が必要
	minMoment := 8000
	count := 0

	// open webcam バックグラウンドで動かすときはここをコメントアウト
	webcam, _ := gocv.VideoCaptureDevice(0)

	//これないと連続撮影できない カメラのindex番号が使用中になるため
	defer webcam.Close()

	// open display window()
	window1 := gocv.NewWindow("detectMotion")

	//差分の画像を表示するウィンドウ
	//window2 := gocv.NewWindow("detectMotion")

	// prepare imag matrix データを格納する準備
	img := gocv.NewMat()
	img1 := gocv.NewMat()
	img2 := gocv.NewMat()
	sabun := gocv.NewMat()

	for {
		// webカメラから画像を取得
		webcam.Read(&img)

		webcam.Read(&img1)

		webcam.Read(&img2)

		// 画像(動画)を表示
		window1.IMShow(img1)
		// これがないとウィンドウが表示されない引数は(1ms待つ)という意味
		window1.WaitKey(1)

		// img1とimg2をグレースケールに変換
		gocv.CvtColor(img1, &img1, gocv.ColorBGRToGray)
		gocv.CvtColor(img2, &img2, gocv.ColorBGRToGray)

		// img1とimg2を2値化
		gocv.Threshold(img1, &img1, 100, 255, gocv.ThresholdBinary)
		gocv.Threshold(img2, &img2, 100, 255, gocv.ThresholdBinary)

		// img1とimg2の差分の画像を作成
		gocv.AbsDiff(img1, img2, &sabun)

		//差分の画像を表示するウィンドウ
		//window2.IMShow(sabun)
		//window2.WaitKey(1)

		//差分で生まれた白色の領域をカウント
		moment := gocv.CountNonZero(sabun)

		//差分(白色の領域)が閾値以上なら動いていると判断
		if moment > minMoment {
			fmt.Printf("moment:%d :異常あり\n", moment)
			count++
		} else {
			fmt.Printf("moment:%d :異常なし\n", moment)
		}

		if count > 20 {
			count = 0
			now := time.Now()
			//img1をimagesフォルダに保存しpathを返す
			path1 := "images/" + now.Format("20060102150405") + ".jpg"
			gocv.IMWrite(path1, img)
			fmt.Println("Lineに通知")
			return path1
		}
	}
}
