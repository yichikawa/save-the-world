package main

import(
	"log"
	"fmt"
	"net/http"
	"html/template"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
)

var password string

func yichikawaHandler(w http.ResponseWriter, r *http.Request) {
	//サーバ起動時に入力したDBのパスワードを含んだログイン情報文字列を生成し、DBに接続
	var passstr = fmt.Sprintf("root:%s@tcp(localhost:3306)/test",password)
	db, err := sql.Open("mysql", passstr )
        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

	//DB接続成否をチェック
        err = db.Ping()
        if err != nil {
                log.Println("データベース接続失敗")
                log.Fatal(err)
        }
	
	//クエリを実行しrowsに結果を取得
	rows, err := db.Query("select name,intro from intro where id=1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	//結果スキャンし変数に代入
	var name string
	var intro string
	rows.Next()
	rows.Scan(&name,&intro)
	
	//テンプレートを使ってデータを含んだhtml？をレスポンスライターに渡す
	t, err :=template.ParseFiles("introtemplate.tpl")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w,intro)	
}	

func rootHandler(w http.ResponseWriter, r *http.Request) {
        //サーバ起動時に入力したDBのパスワードを含んだログイン情報文字列を生成し、DBに接続
        var passstr = fmt.Sprintf("root:%s@tcp(localhost:3306)/test",password)
        db, err := sql.Open("mysql", passstr )
        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        //DB接続成否をチェック
        err = db.Ping()
        if err != nil {
                log.Println("データベース接続失敗")
                log.Fatal(err)
        }

        //クエリを実行しrowsに結果を取得
        rows, err := db.Query("select * from intro")
        defer rows.Close()

        //結果を1行ずつスキャンしレスポンスライターに文字列として出力
        var id int
        var name string
	var intro string
        for rows.Next(){
                err := rows.Scan(&id,&name,&intro)
                if err != nil {
                        log.Fatal(err)
                }
                fmt.Fprintf(w,"ID=%d,名前=%s,紹介=%s\n",id, name, intro)
        }
}

func main(){
	fmt.Println("パスワードを入力")
	fmt.Scan(&password)
	http.HandleFunc("/", rootHandler) // ハンドラを登録
	//ここから下に1行ごとに各員のページに関するハンドラを登録する
		//　（テンプレート）　http.HandleFunc("/英名前/",英名前Handler) //名前
	
	http.HandleFunc("/yichikawa/",yichikawaHandler) //市川
	
	//ハンドラ登録ここまで
	http.ListenAndServe(":8080", nil)
}


