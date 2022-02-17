package tutorial

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

// #3-1 Goルーチーン
// 並行処理（軽量なスレッド）
// 頭にgoをつける
func chapter3_1() {
	go chapter3_1_1("hello")
	chapter3_1_1("world")
	// 必ずしもhello→worldの順にはならない。
}

func chapter3_1_1(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// #3-2 チャネル
// ルーチーン間で値を共有する仕組み
// Channel型の変数を宣言
// ch := make(chan int)
// ch <- v    // v をチャネル ch へ送信する
// v := <-ch  // ch から受信した変数を v へ割り当てる

func chapter3_2() {
	ch := make(chan string)
	// チャネルを引数に渡す
	go chapter3_2_1(ch, "world")

	// 受信しないと次に進まない
	chapter3_1_1("hello")
}

func chapter3_2_1(ch chan string, s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	ch <- "success" //0を送信
}

// #3-3 バッファ
// 容量指定が可能
// 容量以上は入らない
// 受信後は再利用可能
func chapter3_3() {
	ch := make(chan int, 2)
	ch <- 1 // 送信
	ch <- 2 // 送信
	//ch <- 3           //容量がいっぱいで入らない
	fmt.Println(<-ch) // 受信
	ch <- 3           // 空きがあるので入る（FIFO）
	fmt.Println(<-ch) // 受信
	fmt.Println(<-ch)
}

// #3-4 セレクト
// 複数チャネル
// selectで場合分け
func chapter3_4() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go chapter3_4_1(ch1, "")
	//go chapter3_4_1(ch2, "")
	go chapter3_4_1(ch3, "")

	for {
		select {
		case <-ch1:
			fmt.Println("ch1受信しました")
		case <-ch2:
			fmt.Println("ch2受信しました")
		case <-ch3:
			fmt.Println("ch3受信しました")
			return
		default:
			//fmt.Println("何も受信していません。")
		}
	}
}

func chapter3_4_1(ch chan int, s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	ch <- 0 //0を送信
}

// #3-5 カウント
// Goルーチーンの待機
func chapter3_5() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1) //カウントインクリメント
		go func(i int) {
			time.Sleep(2 * time.Second)
			fmt.Println("End:", i)
			wg.Done() //カウントデクリメント
		}(i)
	}

	// カウントが0になるまで待機
	wg.Wait()
	fmt.Println("完了")
}

// #3-6
// ファイル書き込み
func chapter3_6(message string) error {
	// 引数：ファイルパス、フラグ、ファイルパーミッション、
	file, err := os.OpenFile("log.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close() //最後に実行
	fmt.Println(time.Now())
	fmt.Fprintln(file, time.Now().Format("2006/01/02 15:04"), message)
	// 時刻フォーマット
	// 参考：https://qiita.com/icbmuma/items/5617f3fc5bc0215aade2
	// 値は指定されたものを使う

	return nil

}

// #3-7
// ファイル読み込み
func chapter3_7() {
	// 引数：ファイルパス、フラグ、ファイルパーミッション、
	file, err := os.OpenFile("sample2.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() //最後に実行

	bArr, err := ioutil.ReadAll(file)

	fmt.Println("このテキスト作成者は" + string(bArr) + "さんです。")
}
