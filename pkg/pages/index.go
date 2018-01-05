package pages

import (
	"net/http"
	"log"
	"html/template"
)

const (
	SoftTitle = "TPRISIS"
)

func Publish() {
	http.HandleFunc("/", index)

}

//首页数据
type IndexData struct {
	Title string
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/index.html", "web/public/css.html", "web/public/js.html")
	if err != nil {
		log.Println(err)
	}
	index := IndexData{SoftTitle}
	t.Execute(w, index)
}
