package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

var auth string

func homepage(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("homepage.html")
  t.Execute(w, nil)
}

func sign_in(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "http://127.0.0.1:9001/sign_in?client_app=Q-Time&redirect_back_url=http%3A%2F%2Flocalhost%3A9003%2Fcreate_session", http.StatusFound)
}

func create_session(w http.ResponseWriter, r *http.Request) {

  auth_response := r.URL.Query();
  for k, v := range auth_response {

    fmt.Println("k:", k,"\n", "v:", v)

    // Parsing the auth token from the redirect URL given by Q-Auth
    auth_token := v[0]
    fmt.Println(auth_token)

    // Creating a Clinent
    client := &http.Client{}

    // Creating the Request Object
    http_request, err := http.NewRequest("GET", "http://127.0.0.1:9001/api/v1/my_profile", nil)

    // Settin the header and adding auth token to it
    http_request.Header.Set("Authorization", "Token token=" + auth_token)
    http_request.Close = true

    // Sending the request to Q-Auth My Profile API and handling the response
    my_profile_response, err := client.Do(http_request)
    if err != nil {
      log.Fatal(err)
    }
    defer my_profile_response.Body.Close()

    my_profile_details, err := ioutil.ReadAll(my_profile_response.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(my_profile_details)[:255])

    var profile interface{}
    json.Unmarshal(my_profile_details, &profile)
    pro_details := profile.(map[string]interface{})
    fmt.Println(pro_details["data"])

    // details := pro_details["data"]
    // prof_details := details.(map[string]interface{})
    // fmt.Println(prof_details)
    // for value := range pro_details["data"] {

    // fmt.Println("Value:", value)
    // }

  }

  t, _ := template.ParseFiles("create_session.html")
  t.Execute(w, nil)
}

func timesheet(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("timesheet.html")
  t.Execute(w, nil)
}

func sign_out(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "http://127.0.0.1:9003", http.StatusFound)
}

func main() {
    http.HandleFunc("/", homepage)
    http.HandleFunc("/sign_in", sign_in)
    http.HandleFunc("/create_session", create_session)
    http.HandleFunc("/timesheet", timesheet)
    http.HandleFunc("/sign_out", sign_out)
    err := http.ListenAndServe(":9003", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}