package main

import(
        "database/sql"
        "fmt"
	"log"
	"net/http"
        _ "github.com/go-sql-driver/mysql"
)

var password string

func handler(w http.ResponseWriter, r *http.Request) {
	var passstr = fmt.Sprintf("root:%s@tcp(localhost:3306)/test",password)
	log.Printf("%T",passstr)
	db, err := sql.Open("mysql", passstr )
        if err != nil {
                log.Fatal(err)
        }

        defer db.Close()
	err = db.Ping()
	if err != nil {
                log.Println("データベース接続失敗")
                log.Fatal(err)
	}

	rows, err := db.Query("select * from test")
	defer rows.Close()
	
	var id int
	var name string
	for rows.Next(){
		err := rows.Scan(&id,&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w,"ID=%d,Name=%s\n",id, name)
	}
}	

func main(){
	fmt.Println("パスワードを入力")
	fmt.Scan(&password)
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}


