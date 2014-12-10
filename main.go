package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("homepage.html")
  t.Execute(w, nil)
  // fmt.Fprintf(w, "Welcome to Q-Time Homepage!")
}

func sign_in(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("sign_in.html")
  t.Execute(w, nil)
}

func authenticate(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  username := r.Form["username"][0]
  password := r.Form["password"][0]
  fmt.Println("username:", username)
  fmt.Println("password:", password)
  if username == "Kedar" && password == "12345" {
    http.Redirect(w, r, "http://127.0.0.1:3000/timesheet", http.StatusFound)
  } else {
    t, _ := template.ParseFiles("sign_in.html")
    t.Execute(w, nil)
  }
}

func timesheet(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("timesheet.html")
  t.Execute(w, nil)
}

func sign_out(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "http://127.0.0.1:3000/sign_in", http.StatusFound)
}
// func login(w http.ResponseWriter, r *http.Request) {
//     fmt.Println("method:", r.Method)
//     if r.Method == "GET" {
//         t, _ := template.ParseFiles("login.gtpl")
//         t.Execute(w, nil)
//     } else {
//         r.ParseForm()
//         // username := r.Form["username"]
//         // password := r.Form["password"]
//         fmt.Println("username:", r.Form["username"])
//         fmt.Println("password:", r.Form["password"])
//         http.Redirect(w, r, "http://127.0.0.1:3000/show", http.StatusFound)
//         // fmt.Println(show(w, r, username, password))
//     }
// }

// func show(w http.ResponseWriter, r *http.Request) {
//   fmt.Println("method:", r.Method)
//   if r.Method == "GET" {
//     // fmt.Println("username:%s", un)
//     // fmt.Println("password:%s", pass)
//   } else {
//     t, _ := template.ParseFiles("login.gtpl")
//     t.Execute(w, nil)
//   }
// }

func main() {
    http.HandleFunc("/", homepage)
    http.HandleFunc("/sign_in", sign_in)
    // http.HandleFunc("/show", show)
    http.HandleFunc("/authenticate", authenticate)
    http.HandleFunc("/timesheet", timesheet)
    http.HandleFunc("/sign_out", sign_out)
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}