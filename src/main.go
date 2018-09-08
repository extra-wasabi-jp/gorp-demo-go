package main
import (
    "fmt"
    "io/ioutil"
    "database/sql"
//    _ "github.com/lib/pq"
    _ "github.com/mattn/go-sqlite3"
    yaml "gopkg.in/yaml.v2"
    "./repos"
)

var err error
var db *sql.DB
var repo repos.Repo


func main() {

    buf, err1 := ioutil.ReadFile("../config.yml")
    if err1 != nil {
        panic(err)
    }
    var config Config
    err = yaml.Unmarshal(buf, &config)
    if err != nil {
        panic(err)
    }

    connStr := fmt.Sprintf("%s", config.Application.Db.Url)

    fmt.Println("◇ config")
    fmt.Println("---------------------------------------")
    fmt.Printf("%#v\n", config)
    fmt.Println("---------------------------------------")

    db, err = sql.Open("sqlite3", connStr)
    if err != nil {
        fmt.Println("db open error")
        panic(db)
    }
    defer db.Close()

    repo = repos.Repo{Db: db}
    repo.Init()

    fmt.Println("")
    fmt.Println("◇ 社員一覧")
    fmt.Println("---------------------------------------")
    fmt.Printf("ログインID\t氏名（漢字）\t氏名（カナ）\t登録日時\n");
    employeeList := repo.GetEmployeeAll()
    for i := range employeeList {
        emp := employeeList[i]
        fmt.Printf("%s\t%s\t%s\t%s\n",
                      emp.LoginId,
                      emp.UserNmJa,
                      emp.UserNmKa,
                      emp.CreatedAt.Format("2006/01/02 03:04:05"))
    }
    fmt.Println("---------------------------------------")


    fmt.Println("")
    fmt.Println("◇ 鈴木さん")
    fmt.Println("---------------------------------------")
    emp2 := repo.GetEmployeeByLoginId("ichirou.suzuki")
    fmt.Printf("%s\t%s\t%s\t%s\n",
                  emp2.LoginId,
                  emp2.UserNmJa,
                  emp2.UserNmKa,
                  emp2.CreatedAt.Format("2006/01/02 03:04:05"))
    fmt.Println("---------------------------------------")

}
