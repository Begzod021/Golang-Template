package main

import ("fmt"
		"net/http"
		"html/template"
	)


type User struct {
	Username string
	Age uint16
	Money int16
	Avg_grades, Happines float64
	Hobbies []string
}

func (user User) getallinfo() string{
	return fmt.Sprintf("Username is: %s. He is %d and he has money equal: %d", user.Username, user.Age, user.Money)
}

func (user *User) setNewUsername(username string){
	user.Username = username
}

func home_page(w http.ResponseWriter, r *http.Request){
	user := User{"admin", 20, -500, 5, 5.5, []string{"Football", "Chess", "Dance"}}
	tmpl, _ := template.ParseFiles("templates/index.html")
	user.setNewUsername("admin")
	tmpl.Execute(w, user)
}

func contacts_page(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprintf(w, "Contact Page")
}

func handleRequest(){
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8000", nil)
}

func main(){
	handleRequest()
}