package tutorial

import (
	"fmt"
	"strconv"
	"strings"
)

func chapter2() {
	//chapter2_1()
	//chapter2_2()
	//chapter2_2_3()
	//chapter2_3()
	//chapter2_4()
	//chapter2_5_1()
	//chapter2_6()
	//chapter2_7()
	//chapter2_8()
	//chapter2_8_1()
}

// #2-1 文字列操作
func chapter2_1() {
	// よく使う文字列操作
	// 1:部分取得（開始位置：）：[2:]
	//   部分取得（：終了位置）：[:3]
	//   部分取得（開始位置：終了位置）：[2:3]
	// 2:長さ len(string)
	// 3:空白削除 strings.TrimSpace(string) // 入力値、DB値
	// 4:分割 strings.Split(string, "区切り文字") // CSV

	// ↓の文字列からabcdefを抽出せよ
	// "12345abcdefABCDEF"

	// ↓の文字列の長さを求めよ
	// "12345abcdefABCDEF"

	// ↓の空白を削除せよ
	// "   12345abcdefABCDEF   "

	// ↓の文字列をカンマで区切れ
	// "12345,abcdef,ABCDEF"

	fmt.Println("部分取得")
	str := "12345abcdefABCDEF"
	fmt.Println(str[5:11])

	// 長さ len()
	fmt.Println("長さ")
	fmt.Println(len(str))

	// トリム strings.TrimSpace(string)
	fmt.Println("トリム")
	str2 := "   12345abcdefABCDEF   "
	fmt.Println(strings.TrimSpace(str2))

	// 分割 strings.Split(string, "区切り文字")
	fmt.Println("分割")
	str3 := "12345,abcdef,ABCDEF"
	slice := strings.Split(str3, ",")
	for _, s := range slice {
		fmt.Println(s)
	}
	fmt.Println("-------------------------------")
}

// #2-2
// ポインタ：
// ポインタとはデータのアドレスを示す
// メリット：
// 参照渡しができる
func chapter2_2() {
	i := 5
	chapter2_2_1(i) // 値渡し
	fmt.Println(i)  //値は？

	chapter2_2_2(&i) // 参照渡し
	fmt.Println(i)   //値は？
	fmt.Println("-------------------------------")
}

func chapter2_2_1(arg int) {
	arg = 10
}

func chapter2_2_2(arg *int) {
	*arg = 10
}

// intは整数型、*intはポインタ型と変換すれば理解しやすい
// 宣言		var p int 	var p *int
// 値　		p 		  	*p
// アドレス　&p	 		 p

// ↓変数iの値とアドレスを出力せよ
// i := 10

// ↓変数i2に10を代入し、値とアドレスを出力せよ
// var i2 *int

func chapter2_2_3() {
	i := 10
	fmt.Println(i)  //値
	fmt.Println(&i) //アドレス

	var i2 *int
	num := 10
	i2 = &num
	fmt.Println(*i2) //値
	fmt.Println(i2)  //アドレス

	// 余談
	var i3 *int
	*i3 = 10        // エラー
	fmt.Println(i3) // ポインタの初期値はnil

	// invalid memory address or nil pointer dereference←nilが故のエラー
	// fmt.Printlnで出力してみると良い

	var i4 *int
	num2 := 10
	*i4 = num2 // エラー

	fmt.Println("-------------------------------")
}

// #2-3
// 構造体
// type 型の名前 struct {
//     フィールド名 型名
//     フィールド名 型名
//     フィールド名 型名
// }

type User struct {
	userId   string
	userName string
	password string
}

//　宣言
func chapter2_3() {
	// 宣言：変数:=型の名前{}
	u := User{"ID", "ユーザ名", "パスワード"}
	// 値：型.フィールド名
	fmt.Println(u.userId)
	fmt.Println(u.userName)
	fmt.Println(u.password)

	// ポインタを使う場合
	u2 := new(User) //初期化あり
	fmt.Println(u2)
	u2.userId = "ID2"
	u2.userName = "ユーザ名2"
	u2.password = "パスワード2"
	fmt.Println(u2.userId)
	fmt.Println(u2.userName)
	fmt.Println(u2.password)

	u3 := &User{"ID3", "ユーザ名3", "パスワード3"} //初期化あり
	fmt.Println(u3.userId)
	fmt.Println(u3.userName)
	fmt.Println(u3.password)

	var u4 = &User{} //初期化なし
	u4.userId = "ID4"
	u4.userName = "ユーザ名4"
	u4.password = "パスワード4"
	fmt.Println(u4.userId)
	fmt.Println(u4.userName)
	fmt.Println(u4.password)

	var u5 *User
	// fmt.Println(u5)
	u5.userId = "ID5" //エラー
	u5.userName = "ユーザ名5"
	u5.password = "パスワード5"
	fmt.Println(u5.userId)
	fmt.Println(u5.userName)
	fmt.Println(u5.password)
	fmt.Println("-------------------------------")
}

// #2-4
// 効率化コマンド（go fill struct）
// 初期化コードを簡単に作成する

// コマンドを打ってみよう
type Data struct {
	no    int
	name  string
	ret   bool
	data1 string
	data2 string
	data3 string
	data4 string
	data5 string
	data6 string
	data7 string
	data8 string
	data9 string
}

func chapter2_4() {

}

// #2-5
// メソッド
// 構造体専用の関数
func (u User) chapter2_5() {
	// u User：レシーバ
	// レシーバ名はシンプルが推奨
	fmt.Println(u.userId)
	fmt.Println(u.userName)
	fmt.Println(u.password)
}

func chapter2_5() { // 関数名同じ扱いにならない
	u := User{"ID", "ユーザ名", "パスワード"}
	// 呼び出し方法：型の名前.メソッド
	u.chapter2_5()
	fmt.Println("-------------------------------")
}

// ポインタメソッド
func (u *User) chapter2_5_1() {
	u.userId = "ID"
	u.userName = "ユーザ名"
	u.password = "パスワード"
}

func chapter2_5_1() {
	u := &User{}
	// 呼び出し方法：型の名前.メソッド
	u.chapter2_5_1()

	// 参照渡し
	fmt.Println(u.userId)
	fmt.Println(u.userName)
	fmt.Println(u.password)
	fmt.Println("-------------------------------")
}

// #2-6
// インターフェース
// interface：何でもOKな型
func chapter2_6() {
	var i interface{}
	i = 10
	fmt.Println(i)
	i = "str"
	fmt.Println(i)
	i = []string{"arr1", "arr2"}
	fmt.Println(i)
	fmt.Println("-------------------------------")
}

// #2-7
// 型判定
func chapter2_7() {
	var i interface{}
	i = 10
	_, ok := i.(int)
	fmt.Println(ok)
	_, ng := i.(string)
	fmt.Println(ng)

	i = "str"
	_, ok = i.(string)
	fmt.Println(ok)
	_, ng = i.(int)
	fmt.Println(ng)

	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")
	}
	fmt.Println("-------------------------------")
}

// #2-8
// インターフェース（構造体）
// Golangは静的型付け言語のため、変数を使う時には、その変数の型が決まっていないといけない
// インターフェースを使えば型を意識せずに使える
// 処理の共通化

// Human型
type Human struct {
	name   string
	weight int
}

// Humanメソッド
func (h *Human) eat() {
	h.weight = h.weight + 10
	fmt.Printf("%q君の体重は%vkgです。\n", h.name, h.weight)
}

// Animal型
type Animal struct {
	name   string
	weight int
}

// Animalメソッド
func (a *Animal) eat() {
	a.weight = a.weight + 5
	fmt.Printf("%qの体重は%vkgです。\n", a.name, a.weight)
}

// 実行
func chapter2_8() {
	human := Human{"田中くん", 60}
	dog := Animal{"ぽち", 10}

	human.eat() // Human型.eat()
	dog.eat()   // Animal型.eat()
	fmt.Println("-------------------------------")
}

// Human2型
type Human2 struct {
	name   string
	weight int
}

// Human2メソッド
func (h *Human2) eat2() {
	h.weight = h.weight + 10
	fmt.Printf("%q君の体重は%vkgです。\n", h.name, h.weight)
}

// Animal2型
type Animal2 struct {
	name   string
	weight int
}

// Animalメソッド
func (a Animal2) eat2() {
	a.weight = a.weight + 5
	fmt.Printf("%qの体重は%vkgです。\n", a.name, a.weight)
}

// インターフェース
type Common interface {
	eat2() // eat2()のメソッドを持つ構造体はCommon型が使える
}

func chapter2_8_1() {
	human := &Human2{"田中くん", 60}
	dog := &Animal2{"ぽち", 10}

	// インターフェース
	var c Common
	c = human // eat2()を持つ構造体を代入できる
	c.eat2()  //  Common型.eat2()
	c = dog
	c.eat2() //  Common型.eat2()
	fmt.Println("-------------------------------")
}

// #2-9 課題
// Member構造体は以下のフィールドを持つ
// name(string)、id(string)、password(string)、gender(int)、age(int))
// Admin構造体は以下のフィールドを持つ
// name(string)、id(string)、password(string)、note(string)

// 次に、Member構造体とAdmin構造体は以下のメソッドを持つ
// ・情報取得（GetInfo）

// Member、Adminそれぞれ最低2つのオブジェクトを作成し初期化せよ
// IDとパスワードを入力し、マッチするユーザ（MemberもしくはAdmin）が存在する場合、そのユーザの全ての情報を出力するプログラムを作成せよ

const (
	// 性別
	Male   = 0
	Female = 1
)

type Member struct {
	name     string
	id       string
	password string
	gender   int
	age      int
}

type Admin struct {
	name     string
	id       string
	password string
	note     string
}

func Task2_1() {
	// 初期化
	member1 := &Member{
		name:     "浅川",
		id:       "00001",
		password: "pass00001",
		gender:   Male,
		age:      18,
	}
	member2 := &Member{
		name:     "山田",
		id:       "00002",
		password: "pass00002",
		gender:   Female,
		age:      26,
	}
	admin1 := &Admin{
		name:     "管理者",
		id:       "admin1",
		password: "ad00001",
		note:     "東京支店用",
	}
	admin2 := &Admin{
		name:     "管理者2",
		id:       "admin2",
		password: "ad00002",
		note:     "大阪支店用",
	}

	// 入力処理開始
	var input string
	// 身長
	fmt.Println("IDを入力してください。")
	fmt.Scan(&input)
	id := input

	fmt.Println("続けてパスワードを入力してください。")
	fmt.Scan(&input)
	password := input

	var c Change
	switch id + ":" + password {
	case member1.id + ":" + member1.password:
		c = member1
	case member2.id + ":" + member2.password:
		c = member2
	case admin1.id + ":" + admin1.password:
		c = admin1
	case admin2.id + ":" + admin2.password:
		c = admin2
	default:
		fmt.Println("ログインに失敗しました。")
		return
	}

	ret := c.GetInfo()
	show := ""
	for _, r := range ret {
		switch r.(type) {
		case string:
			show = show + r.(string)
		case int:
			show = show + strconv.Itoa(r.(int))
		}
	}

	fmt.Printf("現在のあなたの情報は%qです。\n", show)

}

func (m *Member) GetInfo() (info []interface{}) {
	info = append(info, "名前:")
	info = append(info, m.name)
	info = append(info, "、ID:")
	info = append(info, m.id)
	info = append(info, "、パスワード:")
	info = append(info, m.password)
	info = append(info, "、性別:")
	info = append(info, ConvGender(m.gender))
	info = append(info, "、年齢:")
	info = append(info, m.age)
	return
}

func (a *Admin) GetInfo() (info []interface{}) {
	info = append(info, "、名前:")
	info = append(info, a.name)
	info = append(info, "、ID:")
	info = append(info, a.id)
	info = append(info, "、パスワード:")
	info = append(info, a.password)
	info = append(info, "、備考:")
	info = append(info, a.note)
	return
}

type Change interface {
	GetInfo() []interface{}
}

func ConvGender(i int) string {
	if i == Male {
		return "男性"
	}
	if i == Female {
		return "女性"
	}
	return ""
}
