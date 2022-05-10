package main

import ("fmt"
        "net/http")

type User struct  {
  name string
  age uint16
  money uint16
  avg_grades, hapiness float64
  }

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is: %s. He is %d and he has money " +
    "equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string)  {
  u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request)  {
  bob := User{name: "Bob, 25, -50, 4.2, 0.8"}
  bob.setNewName("Alex")
  fmt.Fprintf(w, "User name is: " + bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Hello World")
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":876", nil)
}

func main()  {
  // var bob User = ....
  // bob := User{name: "Bob, age: 25, money: -50, avg_grades: 4.2, hapiness: 0.8"}
    handleRequest()
}
