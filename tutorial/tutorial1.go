package tutorial

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// #0
// ディレクトリ構成
// $GOPATH（go env）
// |--bin
// |--pkg
// |--src
//    |--github.com
//       |--ユーザ名
//          |--プロジェクト名
//             |--main.go
//             |--go.mod
//             |--go.sum
//             |--パッケージ名
//                |--パッケージ名.go

// #1：main
// mainは特殊関数
// エントリポイント
// package mainのmain関数
func main() {
	// パッケージ名.関数
	//fmt.Println("Hello World")
	// 自動import
	// 実行：go run ファイル名

	//chapter3()

	// if err!=nilでエラーハンドリングするのが基本
	/*
		err := chapter4("o") // 失敗
		if err != nil {
			fmt.Println(err)
		}
		err = chapter4("0") // 成功
		if err != nil {
			fmt.Println(err)
		}*/

	// 複数の戻り値
	/*num, ret := chapter5(10, 5)
	if ret {
		fmt.Println(num)
	}*/

	//chapter6(5, 2, 50)

	//chapter7()

	//chapter8()

	//chapter9()

	//task1()
	//task2()
}

// #2：init
// initはパッケージの初期化関数
// main関数より先に実行される
// 他パッケージのinitはimportされた時に実行される
func init() {
	fmt.Println("init")
}

// #3：変数定義
func chapter3() {
	// 型指定
	var i int
	// 定義した変数は使われないとエラーになる
	fmt.Println(i)

	var i2 = 0
	fmt.Println(i2)

	// var省略
	// 自動で型判定される
	i3 := 0
	fmt.Println(i3)

	s := "文字列"
	fmt.Println(s)

	b := true
	fmt.Println(b)

	f := 0.01
	fmt.Println(f)

	// 初期化しなかった場合の値
	var i4 int
	var f2 float64
	var b2 bool
	var s2 string
	fmt.Printf("%v %v %v %q\n", i4, f2, b2, s2)

	// 一度宣言した後は =
	i5 := 0
	i5 = 9
	fmt.Println(i5)

	// グローバル変数
	fmt.Println(global)
}

var global = "グローバル変数"

// #4：エラーハンドリング
func chapter4(s string) error {
	// 文字列→数値変換
	i, err := strconv.Atoi(s)

	// エラーハンドリング
	if err != nil {
		// error→文字列
		errMessage := err.Error()
		fmt.Println(errMessage)
		// そのままOK
		fmt.Println(err)
		return err
	}
	fmt.Println(i)
	return nil
}

// #5：複数の戻り値
func chapter5(i int, max int) (int, bool) {
	num := rand.Intn(i)
	if num < max {
		return num, false
	}
	return num, true
}

// #6 条件文
func chapter6(i, n, limit float64) { // 型省略
	// if通常
	num := math.Pow(i, n)
	if num <= limit {
		fmt.Printf("%v以下です", limit)
	} else {
		fmt.Printf("%vより上です", limit)
	}
	// 括弧がいらない
	// ↓NG
	if num <= limit {
		fmt.Printf("%v以下です", limit)
	} else {
		fmt.Printf("%vより上です", limit)
	}

	// if省略
	if num2 := math.Pow(i, n); num2 <= limit {
		fmt.Printf("%v以下です", limit)
	} else if num == limit {
		fmt.Printf("%v以下です", limit)
	} else {
		fmt.Printf("%vより上です", limit)
	}

	// select（選択肢が多い時）
	num3 := 2
	switch num3 {
	case 1:
		fmt.Println("値は1です")
	case 2:
		fmt.Println("値は2です")
	case 3:
		fmt.Println("値は3です")
	case 4:
		fmt.Println("値は4です")
	case 5:
		fmt.Println("値は5です")
	default:
		fmt.Println("それ以外です")
	}

	// 処理が同じ場合
	switch num3 {
	case 1, 2, 3, 4, 5:
		fmt.Println("5以下です")
	default:
		fmt.Println("それ以外です")
	}
}

// #7：配列、スライス、map
func chapter7() {
	// 配列
	// 要素数を指定
	arr := [4]string{"str1", "str2", "str3"}
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	// 値の代入
	arr[3] = "str4"
	fmt.Println(arr[3])

	// スライス
	// 要素数不変
	var slice []string
	//fmt.Println(slice[0]) // エラー

	// 要素追加
	slice = append(slice, "str1")
	slice = append(slice, "str2")
	fmt.Println(slice[0]) // 成功
	fmt.Println(slice[1]) // 成功

	// 初期化ver
	slice2 := []string{"str1", "str2", "str3"}
	fmt.Println(slice2[0]) // 成功

	// map
	m := map[string]int{"apple": 150, "banana": 300}
	fmt.Println(m["apple"])  // 150
	fmt.Println(m["banana"]) // 300
	fmt.Println(m["lemon"])  // 存在しない→0 条件に使える

	// 要素の追加
	m2 := map[string]string{}
	m2["apple"] = "りんご"
	fmt.Println(m2["apple"])

	// キーが存在するかの判定
	_, ok := m2["apple"]
	fmt.Println(ok) // true
	_, ok = m2["orange"]
	fmt.Println(ok) // false
}

// #8：for文
func chapter8() {
	var slice []string

	// 基本のfor文
	for i := 0; i < 10; i++ {
		// 数値→文字列変換
		slice = append(slice, "str"+strconv.Itoa(i))
	}

	// 要素の数だけ
	for _, str := range slice {
		// 値
		fmt.Println(str)
	}

	// _で省略できる（今後使わないとき）
	for _, str := range slice {
		fmt.Println(str)
	}

	i := 0
	// 無限ループ
	for {
		i++
		fmt.Println(i)
	}

	// 条件あり
	i = 0
	for i < 10 {
		i++
		fmt.Println(i)
	}
}

// 標準入力
func chapter9() {
	fmt.Println("何か入力してください。")
	var input string
	fmt.Scan(&input)
	fmt.Println(input)
}

// #10 課題
// ユーザが身長と体重を入力したら、BMIと肥満判定を出力するプログラムを作成せよ
// ■参考
// 　BMI ＝ 体重kg ÷ (身長m)2乗
// ■肥満判定
//  〜16.99以下	        痩せ
//  17.00〜18.49以下	痩せぎみ
//  18.50〜24.99以下	普通体重
//  25.00〜29.99以下	前肥満
//  30.00〜以上	        肥満
//

// #11 課題2
// 名前、性別、年齢、出身地を入力し、最もポイントが高い人の名前をおすすめするログラムを作成せよ（全て0ptの場合は該当なし）
// ■ルール
// 性別が同じなら10pt
// 年齢が同じなら10pt
// 年齢が1個差なら9pt
// 年齢が2個差なら8pt
// ・・・・
// 年齢が9個差なら1pt
// 年齢がそれ以外なら0pt
// 出身地が同じなら10pt
// 0ptならおすすめなし

// ■使用する情報（名前,性別,年齢,出身地）
// 田中,男,21歳,北海道
// 竹中,女,19歳,東京
// 渡辺,男,36歳,鹿児島
// 山本,男,23歳,千葉
// 池内,女,40歳,岩手
//
//
//
//
//
//
//
//
//

// メッセージ
var MessageIncorrect = "無効な値です。再度入力してください。"

func task1() {
	// 入力処理開始
	var input string
	// 身長
	fmt.Println("身長(cm)を入力してください。")
	fmt.Scan(&input)
	// 変換
	height, err := strconv.ParseFloat(input, 64)
	for err != nil || height < 0 {
		fmt.Println(MessageIncorrect)
		fmt.Scan(&input)
		height, err = strconv.ParseFloat(input, 64)
	}
	// メートルに変換
	height = height / 100

	// 体重
	fmt.Println("体重(kg)を入力してください。")
	fmt.Scan(&input)
	// 変換
	weight, err := strconv.ParseFloat(input, 64)
	for err != nil || weight < 0 {
		fmt.Println(MessageIncorrect)
		fmt.Scan(&input)
		weight, err = strconv.ParseFloat(input, 64)
	}

	// 計算
	bmi := weight / math.Pow(height, 2)
	fmt.Printf("あなたのBMIは%vです。", bmi)
	if bmi <= 16.99 {
		fmt.Println("あなたは痩せ型です。")

	} else if bmi > 16.99 && bmi <= 18.49 {
		fmt.Println("あなたは痩せぎみです。")

	} else if bmi > 18.49 && bmi <= 24.49 {
		fmt.Println("あなたは普通体重です。")

	} else if bmi > 24.99 && bmi <= 29.99 {
		fmt.Println("あなたは前肥満です。")

	} else {
		fmt.Println("あなたは肥満です。")
	}
}

func task2() {
	// ユーザデータ
	usersInfo := map[int][]string{}
	usersInfo[0] = []string{"田中", "男", "21", "北海道"}
	usersInfo[1] = []string{"竹中", "女", "19", "東京"}
	usersInfo[2] = []string{"渡辺", "男", "36", "鹿児島"}
	usersInfo[3] = []string{"山本", "男", "23", "千葉"}
	usersInfo[4] = []string{"池内", "女", "40", "岩手"}

	// 入力処理開始
	var input string
	// 名前
	fmt.Println("名前を入力してください。")
	fmt.Scan(&input)
	name := input

	// 性別
	fmt.Println("性別を入力してください。")
	fmt.Scan(&input)
	gender := input

	// 年齢
	fmt.Println("年齢を入力してください。")
	fmt.Scan(&input)
	// 数値変換
	old, err := strconv.Atoi(input)
	for err != nil {
		fmt.Println(MessageIncorrect)
		fmt.Scan(&input)
		old, err = strconv.Atoi(input)
	}

	// 出身地
	fmt.Println("出身地を入力してください。")
	fmt.Scan(&input)
	homeTown := input

	// 採点結果
	max := 0 // 最高得点
	pointInfo := map[int]int{}
	// 採点
	for i, info := range usersInfo {
		point := 0
		// 性別
		if gender == info[1] {
			point = point + 10
		}

		// 年齢
		old2, _ := strconv.Atoi(info[2])
		value := math.Abs(float64(old - old2))
		if value < 10 {
			point = point + int(10-value)
		}

		// 出身地
		if homeTown == info[3] {
			point = point + 10
		}

		// 最高得点
		if i == 0 {
			max = point
		} else {
			// 最高値更新
			if max < point {
				max = point
			}
		}
		// 格納
		pointInfo[i] = point
	}

	if max == 0 {
		fmt.Println(name + "さん、あなたにおすすめの人は存在しませんでした。")
		return
	}

	// 結果
	result := ""
	for i, point := range pointInfo {
		if point == max {
			if result == "" {
				result = usersInfo[i][0] + "さん"
			} else {
				result = result + "と" + usersInfo[i][0] + "さん"
			}
		}
	}
	fmt.Println(name + "さん、あなたにおすすめの人は" + result + "です。")
}
