package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	// "encoding/json"
)

// func home(w http.ResponseWriter, r *http.Request){
// 	fmt.Println("We are in home page")
//     http.Handle("/",http.FileServer(http.Dir("./static")))
// }

type Article struct{
	Id int
	Title string 
	Subtitle string
	Content string 
	
}

type Articles []Article

func allarticles(w http.ResponseWriter, r *http.Request){

	// if r.Method == POST
	articles := Articles{
		Article{Id:1,Title:"Test Title",Content:"Hello World",Subtitle:"Test Desc"},
		Article{Id:3,Title:"Title 2",Content:"Hello World2",Subtitle:"Test Desc2"},
	}
	fmt.Println("endpoint 2 reached")
	t,err:= template.ParseFiles("static/index.html")
	fmt.Println(err)
	t.Execute(w,articles[1])
	// json.NewEncoder(w).Encode(articles)


	// name := r.FormValue("name")
	// address := r.FormValue("address")
	// fmt.Fprintf(w, "Name = %s\n", name)
	// fmt.Fprintf(w, "Address = %s\n", address)

}

func main(){
	fmt.Println("Hello World")
	// http.Handle("/",http.FileServer(http.Dir("./static")))
	http.HandleFunc("/articles", allarticles)
	// http.HandleFunc("/articles", allarticles)
	// http.HandleFunc("/articles", allarticles)
	// http.HandleFunc("/articles/search", allarticles)
	// http.HandleFunc("/",home)
	if err:= http.ListenAndServe(":3000", nil);err != nil {
		log.Fatal(err)
	}
}

//POST request for getting the article :  /articles
//GET request for showing all articles :  /articles
//GET request for showing article of particular id  : /articles/<id here>
//GET request for showing article of particular title  : ‘/articles/search?q=<search term here>’