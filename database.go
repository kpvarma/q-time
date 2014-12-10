// package main

// import (
//     "database/sql"
//     "fmt"
//     _ "github.com/bmizerany/pq"
// )

// func main() {
//     db, err := sql.Open("postgres", "user=qtime password=12345 dbname=qtime sslmode=disable")
//     checkErr(err)

//     stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
//     checkErr(err)
//     // username := "qtime"
//     // departname := "IT"
//     // created := "2014-12-09"

//     // res, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES ($1,$2,$3)")
//     // checkErr(err)

//     res, err := stmt.Query("RKReloaded", "IT", "2014-12-09")
//     checkErr(err)

//     id, err := res.LastInsertId()
//     checkErr(err)

//     fmt.Println(id)

//     db.Close()

//   // db, err := sql.Open("postgres", "user=qtime password=password dbname=qtime sslmode=disable")
//   //   checkErr(err)

//   //   //Insert
//   //   stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
//   //   checkErr(err)

//   //   res, err := stmt.Exec("qtime", "IT", "2014-12-09")
//   //   checkErr(err)

//   //   id, err := res.LastInsertId()
//   //   checkErr(err)

//   //   fmt.Println(id)

//   //   // Update
//   //   stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
//   //   checkErr(err)

//   //   res, err = stmt.Exec("update", 1)
//   //   checkErr(err)

//   //   affect, err := res.RowsAffected()
//   //   checkErr(err)

//   //   fmt.Println(affect)

//   //   // Query
//   //   rows, err := db.Query("SELECT * FROM userinfo")
//   //   checkErr(err)

//   //   for rows.Next() {
//   //       var uid int
//   //       var username string
//   //       var department string
//   //       var created string
//   //       err = rows.Scan(&uid, &username, &department, &created)
//   //       checkErr(err)
//   //       fmt.Println(uid)
//   //       fmt.Println(username)
//   //       fmt.Println(department)
//   //       fmt.Println(created)
//   //   }

//   //   // Delete
//   //   stmt, err = db.Prepare("delete from userinfo where uid=$1")
//   //   checkErr(err)

//   //   res, err = stmt.Exec(1)
//   //   checkErr(err)

//   //   affect, err = res.RowsAffected()
//   //   checkErr(err)

//   //   fmt.Println(affect)

//   //   db.Close()
// }

// func checkErr(err error) {
//     if err != nil {
//         panic(err)
//     }
// }

package main

import (
    "fmt"
    "github.com/codegangsta/martini"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
  db, err := sql.Open("postgres", "dbname=qtime sslmode=disable")
  PanicIf(err)
  return db
}

func PanicIf(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {

  m := martini.Classic()
  m.Map(SetupDB())


  m.Get("/", func(w http.ResponseWriter, r *http.Request, db *sql.DB) {

  stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
  PanicIf(err)

  id, err := stmt.Query("Joshi", "Ruby", "2014-12-09")
  PanicIf(err)

  fmt.Println(id)

    rows, err := db.Query("SELECT username,departname,created FROM userinfo")
    PanicIf(err)
    defer rows.Close()

    var username, departname string
    var created string
    for rows.Next() {
      err := rows.Scan(&username, &departname, &created)
      PanicIf(err)
      fmt.Fprintf(w, "Username:%s\n Department Name:%s\n Created At:%s\n", username, departname, created)
    }

  })

  m.Run()

}











