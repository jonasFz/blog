package main

import (
	"fmt"
	"net/http"
	"strings"
)

//Very temporary
func myLog(message string){
	fmt.Println(message);
}

func HandleArticles(w http.ResponseWriter, r *http.Request) {
	if(r.Method != "GET"){
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return;
	}

	index := strings.LastIndex(r.URL.Path, "/")
	articleName := r.URL.Path[index + 1:len(r.URL.Path)]

	//If no name was given then just list the articles wanted
	if articleName == "" {


		articles := getArticleList()

		var sb strings.Builder

		for _, article := range articles{
			sb.WriteString(article.Name)
			sb.WriteString(" ")
			sb.WriteString(article.File)
			sb.WriteString(" ")
			sb.WriteString("/articles/")
			sb.WriteString(article.Name)
			sb.WriteString("\n");
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sb.String()))
	}else{
		fmt.Printf("Will need to serve the file named '%s'\n", articleName)
	}
}

func main(){
	fmt.Println("Test")

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/articles/", HandleArticles)

	http.ListenAndServe("localhost:7991", nil)
}
